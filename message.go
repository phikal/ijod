package main

import (
	"log"
	"time"
)

type Message map[string]interface{}

// Process selects an operation to execute, depending on the value of
// OP.
func (u *User) process(op string, data interface{}) {
	jump := func(d interface{}) {
		if pos, ok := d.(float64); ok {
			cvid.jumpTo(pos, u)
		}
	}

	switch op {
	case "fsync":
		if cvid != nil {
			cvid.fsyncing = true
			send("fsync", nil, u)
		}
	case "ready":
		u.ready = true
		cvid.startp()
	case "join":
		for _, w := range users {
			if u != w {
				w.send("pos", nil, u)
			}
		}
		if len(users) > 1 {
			waiting <- u
		} else {
			u.sendStatus(nil)
		}
		log.Println(u, "joined")
	case "pos":
		if val, ok := data.(float64); ok {
			pos := time.Duration(float64(time.Second) * val)
			u.setPos(pos)
		}
	case "select":
		if name, ok := data.(string); ok {
			selectVideo(name, u)
		}
	case "play":
		cvid.play(u)
	case "pause":
		jump(data)
	}
}
