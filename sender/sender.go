package sender

import "github.com/AntoniusIvan/go-Trial44DockerMultiplePackage/message"

// Sender represents a sender that can send messages
type Sender struct {
	Destination chan *message.Message
}

// NewSender creates a new sender
func NewSender(destination chan *message.Message) *Sender {
	return &Sender{
		Destination: destination,
	}
}

// Send sends a message to the destination
func (s *Sender) Send(message *message.Message) {
	s.Destination <- message
}
