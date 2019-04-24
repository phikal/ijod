package main

import (
	"log"
	"time"
)

type Message struct {
	Name string      `json:"name"`
	From uint32      `json:"from,omitempty"`
	Data interface{} `json:"data"`
}

// Process selects an operation to execute, depending on the value of
// OP.
func (u *User) process(op string, data interface{}) {
	switch op {
	case "ready":
		u.ready = true
		u.room.startp()
	case "join":
		for w := range u.room.users {
			if u != w {
				w.send("pos", nil, u)
			}
		}
		if len(u.room.users) > 1 {
			u.room.wait <- u
		} else {
			u.sendStatus(nil)
		}
		u.room.send("join", u.id, u)
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
