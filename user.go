package main

import (
	"log"
	"sync/atomic"
	"time"
)

var counter uint32

type User struct {
	id    uint32
	msgs  chan Message
	ready bool
	room  *Room
	pos   *time.Duration
}

// newUser adds a new user to room
func newUser(room *Room) *User {
	u := &User{
		msgs: make(chan Message, 1<<4),
		room: room,
	}

	room.Lock()
	room.users[u] = true
	room.Unlock()
	u.id = atomic.AddUint32(&counter, 1)

	return u
}

// send a message to all users
func (r *Room) send(name string, data interface{}, from *User) {
	for u := range r.users {
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
	u.room.Lock()
	delete(u.room.users, u)
	u.room.Unlock()
}

// sendStatus is a meta function to send all necessary information about
// the current state of the room to a user u
func (u *User) sendStatus(pos *time.Duration) {
	if u.room.vid != nil {
		u.send("select", u.room.vid.path, nil)
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
	u.msgs <- Message{
		"name": "list",
		"data": videos,
	}
}

// setPos is used by the client to inform the server of the current
// their position. If they are the last to report their time, the mean
// progress is calculated and sent to the waiting user.
func (u *User) setPos(pos time.Duration) {
	u.room.Lock()
	defer u.room.Unlock()

	u.pos = &pos

	ready := 0
	for w := range u.room.users {
		if w.pos != nil {
			ready++
		}
	}

	if ready+2 >= len(u.room.users) {
		var avg time.Duration
		for u := range u.room.users {
			if u.pos != nil {
				avg += *u.pos
				u.pos = nil
			}
		}
		avg /= time.Duration(len(u.room.users))

		select {
		case user := <-u.room.wait:
			user.sendStatus(&avg)
		case <-time.NewTimer(time.Second * 5).C:
			log.Println("No status received after a second")
		}
	}
}
