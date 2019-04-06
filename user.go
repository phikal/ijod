package main

import (
	"log"
	"sort"
	"sync/atomic"
	"time"
)

var (
	uidc  uint64
	users []*User
)

type User struct {
	id    uint64
	msgs  chan Message
	ready bool
	pos   *time.Duration
}

// newUser adds a new user to room
func newUser() *User {
	u := &User{
		id:   atomic.AddUint64(&uidc, 1),
		msgs: make(chan Message, 1<<4),
	}

	vlock.Lock()
	users = append(users, u)
	vlock.Unlock()
	u.send("msg", "Ijod v1", nil)
	u.send("msg", "<em>Here be dragons...</em>", nil)
	return u
}

//
func send(name string, data interface{}, from *User) {
	for _, u := range users {
		u.send(name, data, from)
	}
}

//
func (u *User) send(name string, data interface{}, from *User) {
	msg := Message{
		"name": name,
		"data": data,
	}
	if from != nil {
		msg["from"] = from.id
	}
	u.msgs <- msg
}

// leave cleans up after a user has closed his connection
func (u *User) leave() {
	vlock.Lock()
	for i, w := range users {
		if w == u {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	vlock.Unlock()
}

//
func (u *User) sendStatus(pos *time.Duration) {
	if pos != nil {
		u.send("time", pos.Seconds(), nil)
	}
	if cvid != nil {
		u.send("select", cvid.path, nil)
	}
	u.send("uid", u.id, nil)
	u.listVideos()
}

// listVideos sends a list of all available videos back to the client
// that requested it
func (u *User) listVideos() {
	var buf []string
	for vid := range videos {
		buf = append(buf, vid)
	}
	sort.Strings(buf)
	u.msgs <- Message{
		"name": "list",
		"data": buf,
	}
}

func (u *User) setPos(pos time.Duration) {
	vlock.Lock()
	defer vlock.Unlock()

	u.pos = &pos

	ready := 0
	for _, w := range users {
		if w.pos != nil {
			ready++
		}
	}

	log.Println(ready, len(users))
	if ready+2 >= len(users) {
		var avg time.Duration
		for _, u := range users {
			if u.pos != nil {
				avg += *u.pos
				u.pos = nil
			}
		}
		avg /= time.Duration(len(users))
		if cvid != nil {
			(<-waiting).sendStatus(&avg)
		}
	}
}
