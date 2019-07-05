package main

import (
	"log"
	"time"
)

// Message contains data sent to or from the server/client
type Message struct {
	Op   string      `json:"op"`
	Name string      `json:"name,omitempty"`
	From int         `json:"from,omitempty"`
	Data interface{} `json:"data"`
}

func (u *User) join() {
	for w := range u.room.users {
		if u != w {
			w.send("pos", nil, u)
		}
	}
	if len(u.room.users) > 1 {
		var users []*User
		for v := range u.room.users {
			if u != v {
				users = append(users, v)
			}
		}
		u.send("who", users, u)
		u.room.wait <- u
	} else {
		u.sendStatus(nil)
	}
	u.room.send("join", u.id, u)

}

// Process selects an operation to execute, depending on the value of
// OP.
func (u *User) process(op string, data interface{}) {
	switch op {
	case "ready":
		u.ready = true
		u.room.startp()
	case "join":
		u.join()
	case "pos":
		if val, ok := data.(float64); ok {
			pos := time.Duration(float64(time.Second) * val)
			u.setPos(pos)
		}
	case "select":
		if name, ok := data.(string); ok {
			err := u.room.selectVideo(name, u)
			if err != nil {
				log.Println(err)
			}
		}
	case "play":
		u.room.play(u)
	case "pause":
		u.room.pause(u)
	case "time":
		if pos, ok := data.(float64); ok {
			u.room.jumpTo(pos, u)
		}
	}
}
