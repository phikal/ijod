package user

import (
	"context"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"ijod/mesg"
	"ijod/room"

	ws "nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var timeout = 10 * time.Second

func Socket(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(ws.StatusInternalError, "Premature disconnect")

	ctx, kill := context.WithCancel(context.Background())
	defer kill()

	id := r.URL.Query().Get("id")
	room, ok := room.GetRoom(id)
	if !ok {
		http.Error(w, "no such room", http.StatusBadRequest)
		return
	}

	user := &mesg.User{
		Name: nextName(),
		Ctx:  ctx,
		Kill: kill,
		In:   make(chan *mesg.Message),
		Out:  make(chan *mesg.Message, 1),
	}
	room.Join(user)
	defer room.Leave(user)

	preselect := r.URL.Query().Get("select")
	if preselect != "" {
		user.Out <- &mesg.Message{
			Type: "state",
			Data: map[string]interface{}{
				"timestamp": time.Now().Format(time.RFC3339),
				"position":  0,
				"video":     `./data/` + preselect,
				"playing":   false,
				"user":      user.Name,
			},
		}

	}

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

			if err = wsjson.Read(ctx, conn, &msg); err != nil {
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
			err := wsjson.Write(ctx, conn, msg)
			if err != nil {
				log.Println(err)
				return
			}
		case <-ctx.Done():
			conn.Close(ws.StatusNormalClosure, "Disconnected")
			log.Println(user.Name, "died")
			return
		}
	}
}
