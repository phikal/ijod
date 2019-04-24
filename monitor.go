package main

import (
	"log"
	"time"
)

func (r *Room) monitor() {
	mon := make(chan *Message)
	r.mon = mon

	for msg := range mon {
		switch msg.Name {
		case "leave":
			go func() {
				time.Sleep(5 * time.Second)
				if len(r.users) == 0 {
					rlock.Lock()
					delete(rooms, r.name)
					rlock.Unlock()
					log.Println("closed room", r.name)
					close(mon)
				}
			}()
		}
	}
}
