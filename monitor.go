package main

import (
	"log"
	"time"
)

func (r *Room) monitor() {
	mon := make(chan *Message)
	r.mon = mon

	for msg := range mon {
		switch msg.Op {
		case "leave":
			go func() {
				time.Sleep(5 * time.Second)
				if len(r.users) == 0 {
					rlock.Lock()
					delete(rooms, r.name)
					rlock.Unlock()
					r.close.Do(func() {
						close(mon)
					})
					log.Println("closed room", r.name)
				}
			}()
		}
	}
}
