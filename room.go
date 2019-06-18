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
	tmpl *template.Template

	rooms = make(map[string]*Room)
	rlock sync.Mutex
)

// Room represents a collection of users and a selected video.
type Room struct {
	sync.Mutex
	name     string
	vid      *Video
	vids     map[string]*Video
	users    map[*User]bool
	hasAdmin bool
	admin    *User
	wait     chan *User
	mon      chan<- *Message
	close    sync.Once
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
<pre>This room has been destroyed. Create a <a href="/">new one</a>.`)
		return
	}

	err := tmpl.Execute(w, struct {
		Name  string
		Admin bool
	}{
		room.name,
		room.hasAdmin,
	})
	if err != nil {
		log.Println(err)
	}
}

func newRoom(admin bool) string {
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
	room.refreshVideos()
	room.hasAdmin = admin

	log.Println("Created room", room.name)
	go room.monitor()
	return "/room?id=" + room.name
}

func (r *Room) refreshVideos() {
	vids := make(map[string]*Video)

	var d map[string]interface{}
	var queue = []map[string]interface{}{videos}
	for len(queue) > 0 {
		d, queue = queue[0], queue[1:]
		for _, i := range d {
			if path, ok := i.(string); ok {
				vids[path] = &Video{path: path}
			} else if m, ok := i.(map[string]interface{}); ok {
				queue = append(queue, m)
			}
		}
	}

	r.Lock()
	r.vids = vids
	r.Unlock()
}

// send a message to all users
func (r *Room) send(name string, data interface{}, from *User) {
	for u := range r.users {
		u.send(name, data, from)
	}
}
