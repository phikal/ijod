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

	user := newUser()
	defer user.leave()

	go func() {
		for msg := range user.msgs {
			// log.Printf("%d sending %v", user.id, msg)
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	var msg Message
	for {
		if err = conn.ReadJSON(&msg); err != nil {
			log.Println(err)
			break
		} else {
			// log.Printf("%d reciving %v", user.id, msg)
			name, ok1 := msg["name"]
			op, ok2 := name.(string)
			if ok1 && ok2 {
				user.process(op, msg["data"])
			}
		}
	}
}
