package p2pNetwork

import (
	"github.com/alexkefer/p2psearch-backend/log"
	"net"
	"sync"
)

// Peer describes metadata corresponding to a peer in the network.
type Peer struct {
	// Addr is the peer's TCP address.
	Addr net.Addr
}

// PeerMap keeps track of all the peers in the peer-to-peer network, including itself.
type PeerMap struct {
	// Mutex should be locked whenever the PeerMap.Peers map is accessed directly. It should not be locked when
	// accessing the PeerMap via its methods.
	Mutex sync.RWMutex
	// Peers maps a Peer struct for every peer in the network, using its address as the key. The PeerMap.Mutex should be
	// locked before accessing PeerMap.Peers directly.
	Peers map[string]Peer
}

// HasPeer checks to see if the PeerMap contains a Peer struct corresponding to the given address.
func (peerMap *PeerMap) HasPeer(addr string) bool {
	peerMap.Mutex.RLock()
	_, hasPeer := peerMap.Peers[addr]
	peerMap.Mutex.RUnlock()
	return hasPeer
}

// AddPeer adds a Peer struct to the PeerMap, using the string form of Peer.Addr as its key.
func (peerMap *PeerMap) AddPeer(peer Peer) {
	log.Debug("adding peer: %s", peer.Addr.String())

	peerMap.Mutex.Lock()
	key := peer.Addr.String()
	peerMap.Peers[key] = peer
	peerMap.Mutex.Unlock()

}

// RemovePeer removes a Peer from the PeerMap, using the given address to determine what Peer to remove.
func (peerMap *PeerMap) RemovePeer(addr string) {

	if !peerMap.HasPeer(addr) {
		log.Warn("cannot remove peer that does not exist: %s", addr)
		return
	}

	log.Debug("removing peer: %s", addr)

	peerMap.Mutex.Lock()
	delete(peerMap.Peers, addr)
	peerMap.Mutex.Unlock()

}
