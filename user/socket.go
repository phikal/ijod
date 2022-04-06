package user

import (
	"log"
	"net/http"

	"ijod/mesg"
	"ijod/room"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

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

	user := &mesg.User{
		Name: nextName(),
		In:   make(chan *mesg.Message),
		Out:  make(chan *mesg.Message),
	}
	room.Join(user)
	defer room.Leave(user)

	go func() {
		for {
			var msg mesg.Message
			if err = conn.ReadJSON(&msg); err != nil {
				log.Println(err)
				break
			}

			log.Printf("[in] %s: %#v", user.Name, msg)
			user.In <- &msg
		}
		close(user.Out)
	}()

	for msg := range user.Out {
		log.Printf("[out] %s: %#v", user.Name, *msg)
		err := conn.WriteJSON(*msg)
		if err != nil {
			log.Println(err)
		}
	}
}
