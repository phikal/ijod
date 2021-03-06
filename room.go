package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	tmpl *template.Template

	rooms = make(map[string]*Room)
	rlock sync.Mutex

	useWiki string
)

// Room represents a collection of users and a selected video.
type Room struct {
	sync.Mutex
	name     string
	vid      *Video
	users    map[*User]bool
	wait     chan *User
	mon      chan<- *Message
	close    sync.Once
	playing  bool
	updated  time.Time
}

func init() {
	index, err := Asset("room.html")
	if err != nil {
		log.Fatalln(err)
	}
	tmpl = template.Must(template.New("room").Parse(string(index)))

	rand.Seed(time.Now().Unix())
}

func room(w http.ResponseWriter, r *http.Request) {
	room, ok := rooms[r.URL.Query().Get("id")]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html>
<title>Ijod?</title>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width" />
<pre>This room has been destroyed. Create a <a href="/new">new one</a>.`)
		return
	}

	err := tmpl.Execute(w, struct {
		Name    string
		UseWiki string
	}{
		room.name,
		strings.TrimSuffix(useWiki, "/"),
	})
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
	go room.monitor()
	return "/room?id=" + room.name
}

// send a message to all users
func (r *Room) send(name string, data interface{}, from *User) {
	for u := range r.users {
		u.send(name, data, from)
	}
}
