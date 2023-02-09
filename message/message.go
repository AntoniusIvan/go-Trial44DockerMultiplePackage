package message

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func TryMainPkg(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal("HellowIVaunMultiple")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	//w.Write([]byte("oioi"))
	w.Write([]byte(string(jsonData)))
}
