package mesg

import "encoding/json"

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
	From *User       `json:"-"`
}

type User struct {
	Name string
	In   chan *Message
	Out  chan *Message
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Name)
}
