package httpServer

import (
	"fmt"
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"github.com/alexkefer/p2psearch-backend/utils"
	"github.com/alexkefer/p2psearch-backend/webDownloader"
	"html"
	"net"
	"net/http"
	"net/url"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "bad request")
}

func shutdownHandler(w http.ResponseWriter, r *http.Request, shutdownChan chan<- bool) {
	shutdownChan <- true
	fmt.Fprintf(w, "shutting down server...")
}

func peersHandler(w http.ResponseWriter, r *http.Request, peerMap *p2pNetwork.PeerMap) {
	peerMap.Mutex.RLock()
	for key, _ := range peerMap.Peers {
		fmt.Fprintf(w, "%s\n", html.EscapeString(key))
	}
	peerMap.Mutex.RUnlock()
}

func cacheFileHandler(w http.ResponseWriter, r *http.Request, fileDataStore *fileData.FileDataStore) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Warn("error caching file: %s", err)
		fmt.Fprintf(w, "error caching file: %s", err)
	} else {
		err = webDownloader.CacheResource(path, fileDataStore)
		if err == nil {
			fmt.Fprintf(w, "server cached resource found at %q", path)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "server failed to cache resource: %s", err)
		}
	}
}

func retrieveFileHandler(w http.ResponseWriter, r *http.Request, fileData *fileData.FileDataStore, peerMap *p2pNetwork.PeerMap, myAddr net.Addr) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error retrieving file: %s", err)
		return
	}

	filename := utils.UrlToFilename(path)

	if fileData.HasFileStored(filename) {
		metadata := fileData.RetrieveFileData(filename)

		file, openErr := os.Open(metadata.FileLoc)

		if openErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "server couldn't open file for: %q", filename)
			log.Error("couldn't open file %q: %s", metadata.FileLoc, openErr)
			return
		}

		w.Header().Set("Content-Type", metadata.FileType)
		_, fileErr := file.WriteTo(w)

		if fileErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("couldn't read file %q: %s", metadata.FileLoc, fileErr)
			return
		}

		log.Debug("retrieving file: %s, mime type: %s", metadata.FileLoc, metadata.FileType)
	} else {
		log.Info("server had no file for %s", filename)

		// Ask other hosts on the network if they have it.
		channel := make(chan *p2pNetwork.Message)
		counter := 0
		peerMap.Mutex.RLock()
		for key, peer := range peerMap.Peers {
			if peer.Addr != myAddr {
				counter += 1
				go func() {
					log.Debug("asking %s for file: %s", key, filename)
					message, fileReqErr := p2pNetwork.SendFileRequest(peer.Addr, myAddr, filename)
					if fileReqErr == nil {
						channel <- message
					} else {
						channel <- nil
					}
				}()
			}
		}
		peerMap.Mutex.RUnlock()

		for i := 0; i < counter; i++ {
			message := <-channel
			if message != nil {
				log.Debug("got %d response from %s", message.Code, message.SenderAddr)
				if message.Code == p2pNetwork.HasFileResponse {
					log.Debug("got %s from peer", filename)
					webDownloader.SaveFile(message.Data, filename, message.DataType, fileData)
				}
			}
		}

		if fileData.HasFileStored(filename) {
			metadata := fileData.RetrieveFileData(filename)

			file, openErr := os.Open(metadata.FileLoc)

			if openErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "server couldn't open file for: %q", filename)
				log.Error("couldn't open file %q: %s", metadata.FileLoc, openErr)
				return
			}

			_, fileErr := file.WriteTo(w)

			if fileErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Error("couldn't read file %q: %s", metadata.FileLoc, fileErr)
				return
			}

			w.Header().Add("Content-Type", metadata.FileType)
		} else {
			w.WriteHeader(http.StatusNotFound)
			log.Warn("couldn't find file on network: %s", filename)
			fmt.Fprintf(w, "network has no resource associated with path: %q", filename)
			return
		}
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request, myAddr net.Addr, peerMap *p2pNetwork.PeerMap) {
	path, err := getPathParam(r.URL)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", err)
		log.Warn("failed to connect to peer: %s", err)
		return
	}

	targetAddr, addrParseErr := net.ResolveTCPAddr("tcp", path)

	if addrParseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", addrParseErr)
		log.Warn("failed to connect to peer: %s", addrParseErr)
		return
	}

	connectErr := p2pNetwork.Connect(myAddr, targetAddr, peerMap)

	if connectErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "failed to connecting to peer: %s", connectErr)
		return
	}

	fmt.Fprintf(w, "sent add me request to: %s", targetAddr.String())
	log.Info("sent add me request to: %s", targetAddr.String())
}

func disconnectHandler(w http.ResponseWriter, myAddr net.Addr, peerMap *p2pNetwork.PeerMap) {
	log.Info("received http disconnect request")
	p2pNetwork.Disconnect(myAddr, peerMap)
	fmt.Fprintf(w, "disconnecting from p2p network")
}

func hostnameHandler(w http.ResponseWriter) {
	name, err := os.Hostname()

	if err == nil {
		fmt.Fprintf(w, "%s", name)
	} else {
		log.Warn("couldn't get hostname: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getPathParam(fromUrl *url.URL) (string, error) {
	params, err := url.ParseQuery(fromUrl.RawQuery)

	if err != nil {
		return "", err
	}

	if params.Has("path") {
		return params.Get("path"), nil
	}

	return "", fmt.Errorf("request has no path param")
}
