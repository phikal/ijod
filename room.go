package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	T *template.Template

	rooms = make(map[string]*Room)
	rlock sync.Mutex
)

type Room struct {
	sync.Mutex
	name  string
	vid   *Video
	users map[*User]bool
	wait  chan *User
}

func init() {
	index, err := Asset("room.html")
	if err != nil {
		log.Fatalln(err)
	}
	T = template.Must(template.New("room").Parse(string(index)))

	rand.Seed(time.Now().Unix())
}

func room(w http.ResponseWriter, r *http.Request) {
	room, ok := rooms[r.URL.Query().Get("id")]
	if !ok {
		http.Error(w, "no such room", http.StatusBadRequest)
		return
	}

	err := T.Execute(w, room.name)
	if err != nil {
		log.Println(err)
	}
}

func newRoom() string {
	rlock.Lock()
	defer rlock.Unlock()

	var room *Room
	for {
		r := rand.Uint64()%(1<<(4*5)) + (1 << (4 * 4))
		name := fmt.Sprintf("%x", r)
		if _, ok := rooms[name]; !ok {
			room = &Room{
				name:  name,
				users: make(map[*User]bool),
				wait:  make(chan *User),
			}
			rooms[name] = room
			break
		}
	}

	log.Println("Created room", room.name)
	return "/room?id=" + room.name
}
