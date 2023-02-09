package receiver

import "github.com/AntoniusIvan/go-Trial44DockerMultiplePackage/message"

// Receiver represents a receiver that can receive messages
type Receiver struct {
	Source chan *message.Message
}

// NewReceiver creates a new receiver
func NewReceiver(source chan *message.Message) *Receiver {
	return &Receiver{
		Source: source,
	}
}

// Receive receives a message from the source
func (r *Receiver) Receive() *message.Message {
	return <-r.Source
}
