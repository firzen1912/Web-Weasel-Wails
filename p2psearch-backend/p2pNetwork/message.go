package p2pNetwork

import (
	"encoding/gob"
	"github.com/alexkefer/p2psearch-backend/log"
	"io"
)

// Message describes a message that can be used to communicate with other peers on the peer-to-peer network.
type Message struct {
	// Code is an int value which describes the kind of message.
	Code int
	// SenderAddr is the string form of the senders TCP peer-to-peer address.
	SenderAddr string
	// DataType is used when sending a HasFileResponse message to describe the resource's MIME type. Otherwise, it can
	// be left as nil.
	DataType string
	// Peers is an array of strings, used when sending a SharePeersResponse message. It contains the string
	// representation of every peer's TCP address. For other message codes it should be left as nil.
	Peers []string
	// Data is used when a message contains some data that needs to be sent as well. HasFileResponse messages send the
	// resource file in the field.
	Data []byte
}

const (
	AddMeRequest = iota
	SharePeersRequest
	SharePeersResponse
	BroadcastMessage
	RemoveMeRequest
	FileRequest
	HasFileResponse
	NoFileResponse
)

// ReceiveMessage decodes a Message struct from an io.Reader, encoded using the "encoding/gob" module.
func ReceiveMessage(r io.Reader) (Message, error) {
	decoder := gob.NewDecoder(r)

	var message Message

	err := decoder.Decode(&message)
	if err != nil {
		log.Error("failed receiving p2p message: %s", err)
	}
	return message, err
}

// SendMessage encodes and sends a Message struct to the given io.Writer. It encodes the message using the
// "encoding/gob" module.
func SendMessage(w io.Writer, message Message) error {
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(message)

	if err != nil {
		log.Error("failed sending p2p message: %s", err)
	}

	return err
}
