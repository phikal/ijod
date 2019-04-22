package main

import (
	"log"
	"time"
)

type Message map[string]interface{}

// Process selects an operation to execute, depending on the value of
// OP.
func (u *User) process(op string, data interface{}) {
	switch op {
	case "ready":
		u.ready = true
		cvid.startp()
	case "join":
		for w := range users {
			if u != w {
				w.send("pos", nil, u)
			}
		}
		if len(users) > 1 {
			waiting <- u
		} else {
			u.sendStatus(nil)
		}
		send("join", u.id, u)
	case "pos":
		if val, ok := data.(float64); ok {
			pos := time.Duration(float64(time.Second) * val)
			u.setPos(pos)
		}
	case "select":
		if name, ok := data.(string); ok {
			err := selectVideo(name, u)
			if err != nil {
				log.Println(err)
			}
		}
	case "play":
		cvid.play(u)
	case "pause":
		if pos, ok := data.(float64); ok {
			cvid.jumpTo(pos, u)
		}
	}
}
