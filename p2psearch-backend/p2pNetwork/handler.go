package p2pNetwork

import (
	"github.com/alexkefer/p2psearch-backend/fileData"
	"github.com/alexkefer/p2psearch-backend/log"
	"net"
	"os"
)

// StartServer starts the peer-to-peer TCP server. Should pass own TCP address, initialized PeerMap and initialized
// fileData.FileDataStore. This function does not return.
func StartServer(myAddr net.Addr, peerMap *PeerMap, files *fileData.FileDataStore) {
	listener, listenErr := net.Listen("tcp", myAddr.String())

	if listenErr != nil {
		log.Error("tcp listen error: %s", listenErr)
		return
	}

	// TODO: The server should bee stopped when the program is exiting.
	for {
		conn, connErr := listener.Accept()

		if connErr != nil {
			log.Error("tcp connection error: %s", connErr)
		} else {
			// Here we create a separate goroutine (thread) to handle this connection
			go HandleConnection(myAddr, conn, peerMap, files)
		}
	}
}

// HandleConnection handles TCP connections the peer-to-peer server receives from other peers. It decodes a Message
// struct from the connection and responds based on the Message.Code field. Should pass own net.Addr, net.Conn struct
// from listener.Accept() after checking for errors, initialized PeerMap and initialized fileData.FileDataStore.
func HandleConnection(myAddr net.Addr, conn net.Conn, peerMap *PeerMap, files *fileData.FileDataStore) {
	message, messageErr := ReceiveMessage(conn)

	if messageErr != nil {
		log.Error("message receive error: %s", messageErr)
		return
	}

	switch message.Code {

	case AddMeRequest:
		// Peer is asking us to add them to our p2pServer map
		addrStr := message.SenderAddr

		peerAddr, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			log.Error("addr parse error: %s", addrParseErr)
			return
		}

		peer := Peer{Addr: peerAddr}
		peerMap.AddPeer(peer)

	case SharePeersRequest:
		// Peer is asking us for our p2pServer map

		addrStr := message.SenderAddr
		_, addrParseErr := net.ResolveTCPAddr("tcp", addrStr)

		if addrParseErr != nil {
			log.Error("addr parse error: %s", addrParseErr)
			CloseConnection(conn)
			return
		}

		peerSlice := make([]string, 0)

		for addr, _ := range peerMap.Peers {
			peerSlice = append(peerSlice, addr)
		}

		peerMap.Mutex.RLock()
		message = Message{Code: SharePeersResponse, Peers: peerSlice, SenderAddr: myAddr.String()}
		_ = SendMessage(conn, message)
		peerMap.Mutex.RUnlock()

	case BroadcastMessage:
		log.Info("received broadcast message from %s: %s", message.SenderAddr, message.Data)

	case RemoveMeRequest:
		peerMap.RemovePeer(message.SenderAddr)

	case FileRequest:
		path := string(message.Data)
		log.Debug("peer %s is asking for file: %s", message.SenderAddr, path)

		if files.HasFileStored(path) {
			metadata := files.RetrieveFileData(path)
			data, readErr := os.ReadFile(metadata.FileLoc)

			if readErr != nil {
				log.Error("could not read file: %s", readErr)
				message = Message{Code: NoFileResponse, SenderAddr: myAddr.String()}
			} else {
				log.Debug("found requested file: %s", path)
				message = Message{Code: HasFileResponse, SenderAddr: myAddr.String(), Data: data, DataType: metadata.FileType}
			}

		} else {
			log.Warn("couldn't find requested file: %s", path)
			message = Message{Code: NoFileResponse, SenderAddr: myAddr.String()}
		}

		sendErr := SendMessage(conn, message)

		if sendErr != nil {
			log.Warn("send message error: %s", sendErr)
		}

	default:
		log.Warn("invalid code %d, closing connection", message.Code)

	}

	CloseConnection(conn)
}

// CloseConnection closes net.Conn struct and reports any errors raised during closing to log.
func CloseConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Error("connection close error: %s", err)
	}
}
