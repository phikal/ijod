package room

import (
	"log"
	"time"

	"ijod/mesg"
	"ijod/tree"
)

const timeout = 10 * time.Minute

func (r *Room) daemon() {
	var (
		users   = make(map[*mesg.User]struct{})
		current interface{}

		// Send a message to everyone interested
		broadcast = func(msg *mesg.Message) {
			log.Printf("Broadcasting(%s): %#v", r.Name, *msg)
			for u := range users {
				if u == msg.From {
					continue
				}
				u.Out <- msg
			}
		}

		// Broadcast a list of the present users
		attending = func() {
			var us []*mesg.User
			for u := range users {
				us = append(us, u)
			}
			broadcast(&mesg.Message{
				Type: "users",
				Data: us,
			})
		}

		// Merge incoming messages into one channel
		mux    = make(chan *mesg.Message)
		manage = func(u *mesg.User) {
			for {
				select {
				case msg := <-u.In:
					msg.From = u
					log.Printf("Received %#v from %s@%s",
						*msg, u.Name, r.Name)
					mux <- msg
				case <-u.Ctx.Done():
					log.Println("Kill", u.Name, "in", r.Name)
					r.leave <- u
					return
				}
			}
		}

		tick = time.NewTicker(timeout)
	)
	defer func() {
		for user := range users {
			user.Kill()
		}
		tick.Stop()
		r.Forget()
	}()

	for {
		tick.Reset(timeout)
		select {
		case user := <-r.enter:
			log.Println(user.Name, "joined", r.Name)
			users[user] = struct{}{}
			go manage(user)
			user.Out <- &mesg.Message{
				Type: "self",
				Data: user.Name,
			}
			attending()
			user.Out <- tree.Message()
			if current != nil {
				user.Out <- &mesg.Message{
					Type: "state",
					Data: current,
				}
			}
		case user := <-r.leave:
			log.Println(user.Name, "left", r.Name)
			delete(users, user)
			attending()
		case msg := <-mux:
			switch msg.Type {
			case "state":
				current = msg.Data
				broadcast(msg)
			case "refresh":
				msg.From.Out <- tree.Message()
			}
		case <-tick.C:
			if len(users) == 0 {
				return
			}
		}
	}
}
