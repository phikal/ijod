package room

import (
	"log"
	"path"
	"time"

	"ijod/mesg"
	"ijod/tree"
	"ijod/ytdl"
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

		// Handle download requests
		download = func(uri, user string) {
			file, err := ytdl.Download(uri)
			if err != nil {
				log.Print(err)
				return
			}
			// XXX: race condition possible
			r.files = append(r.files, file)

			video := path.Join("dl", path.Base(file))
			mux <- &mesg.Message{
				Type: "state",
				Data: struct {
					Ti string  `json:"timestamp"`
					Po float64 `json:"position"`
					Pl bool    `json:"playing"`
					Vi string  `json:"video"`
					Us string  `json:"user"`
				}{
					time.Now().Format(time.RFC3339),
					0, false, video, user,
				},
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
			case "download":
				if uri, ok := msg.Data.(string); ok {
					go download(uri, msg.From.Name)
				}
			}
		case <-tick.C:
			if len(users) == 0 {
				return
			}
		}
	}
}
