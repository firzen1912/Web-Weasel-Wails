// backend package for the peer-to-peer web cache backend. This package contains the entry point for the backend.
package backend

import (
	"Web-Weasel/backend/fileData"
	"Web-Weasel/backend/httpServer"
	"Web-Weasel/backend/log"
	"Web-Weasel/backend/p2pNetwork"
	"Web-Weasel/backend/utils"
	"flag"
	"net"
	"os"
	"os/signal"
)

// run starts the peer-to-peer backend and http web server.
func Run() {
	var connect string
	var p2pPort string
	var httpPort string

	flag.StringVar(&connect, "connect", "", "Peer IP address the backend should connect to initially")
	flag.StringVar(&p2pPort, "p2p-port", "", "Port the p2p server should use")
	flag.StringVar(&httpPort, "http-port", "", "Port the http server should use")

	flag.Parse()

	if p2pPort == "" {
		var portErr error
		p2pPort, portErr = utils.FindOpenPort(9000, 9100)
		if portErr != nil {
			log.Error("error finding open port: %s", portErr)
			return
		}
	} else {
		p2pPort = ":" + p2pPort
	}

	address := utils.GetLocalIPAddress() + p2pPort
	myP2PAddr, myAddrParseErr := net.ResolveTCPAddr("tcp", address)

	if myAddrParseErr != nil {
		log.Error("couldn't parsing my ip arg: %s", myAddrParseErr)
		return
	}

	peerMap := p2pNetwork.PeerMap{Peers: make(map[string]p2pNetwork.Peer)}
	myPeer := p2pNetwork.Peer{Addr: myP2PAddr}
	peerMap.AddPeer(myPeer)

	fileDataStore := fileData.CreateFileDataStore()
	fileDataStore.LoadFileData()

	go p2pNetwork.StartServer(myP2PAddr, &peerMap, &fileDataStore)
	log.Info("my p2p server address: %s", myP2PAddr)

	// If an address is given, try to join its network
	if connect != "" {
		seedAddr, addrParseErr := net.ResolveTCPAddr("tcp", connect)

		if addrParseErr != nil {
			log.Error("seedAddr parse error: %s", addrParseErr)
			return
		} else {
			connectErr := p2pNetwork.Connect(myP2PAddr, seedAddr, &peerMap)

			if connectErr != nil {
				return
			}
		}
	}

	exitChannel := make(chan bool)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	go func() {
		for range osSignals {
			exitChannel <- true
		}
	}()

	go httpServer.StartServer(&peerMap, &fileDataStore, exitChannel, myP2PAddr, httpPort)

	for {
		if <-exitChannel {
			break
		}
	}

	p2pNetwork.Disconnect(myP2PAddr, &peerMap)
	fileDataStore.SaveFileDataStore()
}
