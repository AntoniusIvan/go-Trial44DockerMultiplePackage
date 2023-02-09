package message

// Message represents a message with a text and a timestamp
type Message struct {
	Text string
	Time int64
}

// NewMessage creates a new message with a given text
func NewMessage(text string) *Message {
	return &Message{
		Text: text,
		//Time: lorem.Time(),
	}
}
