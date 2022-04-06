package user

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"ijod/mesg"
	"ijod/room"

	ws "github.com/gorilla/websocket"
)

var (
	upgrader = ws.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
	}
	timeout = 10 * time.Second
)

func Socket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := r.URL.Query().Get("id")
	room, ok := room.GetRoom(id)
	if !ok {
		http.Error(w, "no such room", http.StatusBadRequest)
		return
	}

	ctx, kill := context.WithCancel(context.Background())
	user := &mesg.User{
		Name: nextName(),
		Ctx:  ctx,
		Kill: kill,
		In:   make(chan *mesg.Message),
		Out:  make(chan *mesg.Message),
	}
	room.Join(user)
	defer room.Leave(user)

	var ping int32
	go func() {
		time.Sleep(timeout)
		for atomic.LoadInt32(&ping) == 0 {
			user.Out <- &mesg.Message{Type: "ping"}
			atomic.AddInt32(&ping, 1)
			time.Sleep(timeout)
		}

		log.Println(user.Name, "detected timeout")
		kill()
	}()

	go func() {
		for {
			var msg mesg.Message
			if err = conn.ReadJSON(&msg); err != nil {
				log.Println(err)
				break
			}
			if msg.Type == "pong" {
				if atomic.SwapInt32(&ping, 0) == 0 {
					log.Println(user.Name, "preemptive pong")
					kill()
					return
				}
				continue
			}

			log.Printf("[in] %s: %#v", user.Name, msg)
			user.In <- &msg
		}
	}()

	for {
		select {
		case msg := <-user.Out:
			log.Printf("[out] %s: %#v", user.Name, *msg)
			err := conn.WriteJSON(*msg)
			if err != nil {
				log.Println(err)
				return
			}
		case <-ctx.Done():
			log.Println(user.Name, "died")
			return
		}
	}
}
