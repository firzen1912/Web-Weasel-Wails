// Package httpServer implements a RESTful HTTP server used for serving cached resources to the browser.
//
// This package contains the StartServer routine for starting the HTTP server, as well as handler routines for the
// various API calls HTTP clients can make to the server.
package httpServer

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"github.com/alexkefer/p2psearch-backend/p2pNetwork"
	"github.com/alexkefer/p2psearch-backend/utils"
	"net"
	"net/http"
)

// StartServer starts an HTTP sever used to serve cached resources to the browser. Should pass references to an
// initialized p2pNetwork.PeerMap, an initialized fileData.FileDataStore, a boolean channel transmitter, and a net.Addr
// representing the server's address. The shutdownChan argument will receive true when an HTTP client requests the
// server shutdown. The HTTP server will mutate the peerMap and fileDataStore arguments at the request of an HTTP
// client. This function will not return.
func StartServer(peerMap *p2pNetwork.PeerMap, fileDataStore *fileData.FileDataStore, shutdownChan chan<- bool, myAddr net.Addr, port string) {
	http.HandleFunc("/", defaultHandler)

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		shutdownHandler(w, r, shutdownChan)
	})

	http.HandleFunc("/peers", func(w http.ResponseWriter, r *http.Request) {
		peersHandler(w, r, peerMap)
	})

	http.HandleFunc("/cache", func(w http.ResponseWriter, r *http.Request) {
		cacheFileHandler(w, r, fileDataStore)
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		retrieveFileHandler(w, r, fileDataStore, peerMap, myAddr)
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		connectHandler(w, r, myAddr, peerMap)
	})

	http.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		disconnectHandler(w, myAddr, peerMap)
	})

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		hostnameHandler(w)
	})

	if port == "" {
		var err error
		port, err = utils.FindOpenPort(8080, 8180)

		if err != nil {
			log.Error("could not find open port for http server: %s", err)
			return
		}
	} else {
		port = ":" + port
	}

	log.Info("opening http server on port %s", port)

	http.ListenAndServe(port, nil)
}
