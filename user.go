package main

import (
	"log"
	"sort"
	"time"
)

var (
	uidc  uint
	users = make(map[*User]bool)
)

type User struct {
	id    uint
	msgs  chan Message
	ready bool
	pos   *time.Duration
}

// newUser adds a new user to room
func newUser() *User {
	u := &User{msgs: make(chan Message, 1<<4)}

	vlock.Lock()
	users[u] = true
	uidc += 1
	vlock.Unlock()
	u.id = uidc

	u.send("msg", "Ijod v1", nil)
	u.send("msg", "<em>Here be dragons...</em>", nil)
	return u
}

// send a message to all users
func send(name string, data interface{}, from *User) {
	for u := range users {
		u.send(name, data, from)
	}
}

// send a message to a specific user
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
	delete(users, u)
	vlock.Unlock()
}

//
func (u *User) sendStatus(pos *time.Duration) {
	if cvid != nil {
		u.send("select", cvid.path, nil)
	}
	if pos != nil {
		u.send("time", pos.Seconds(), nil)
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
	for w := range users {
		if w.pos != nil {
			ready++
		}
	}

	if ready+2 >= len(users) {
		var avg time.Duration
		for u := range users {
			if u.pos != nil {
				avg += *u.pos
				u.pos = nil
			}
		}
		avg /= time.Duration(len(users))

		if cvid != nil {
			select {
			case user := <-waiting:
				user.sendStatus(&avg)
			case <-time.NewTimer(time.Second).C:
				log.Println("No status received after a second")
			}
		} else {
			select {
			case user := <-waiting:
				user.sendStatus(&avg)
			case <-time.NewTimer(time.Second).C:
				log.Println("No status received after a second")
			}
		}
	}
}
