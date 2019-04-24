package main

import (
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// socket initialized the communication between a client and the server
// using a websocket protocol
func socket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id")
	room, ok := rooms[id]
	if !ok {
		log.Println("no such room", id)
		http.Error(w, "no such room", http.StatusBadRequest)
		return
	}

	user := newUser(room)
	defer user.leave()

	go func() {
		for msg := range user.msgs {
			log.Println(msg)
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		var msg Message
		if err = conn.ReadJSON(&msg); err != nil {
			log.Println(err)
			break
		} else {
			user.process(msg.Name, msg.Data)
			room.mon <- &msg
		}
	}
}
