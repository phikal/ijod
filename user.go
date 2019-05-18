package main

import (
	"encoding/json"
	"log"
	"time"
)

var counter int

// User stores all information related to a user in a room and it's
// websocket connection
type User struct {
	id    int
	name  string
	msgs  chan *Message
	ready bool
	room  *Room
	pos   *time.Duration
}

// MarshalJSON transforms a User into a JSON object, by turning it into
// it's name-value as a string
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.name)
}

// newUser adds a new user to room
func newUser(room *Room) *User {
	room.Lock()
	defer room.Unlock()
	u := &User{
		id:   counter,
		msgs: make(chan *Message, 1<<4),
		room: room,
		name: words[counter%len(words)],
	}

	room.users[u] = true
	u.send("uid", u.id, u)
	counter++

	return u
}

// send a message to a specific user
func (u *User) send(op string, data interface{}, from *User) {
	msg := Message{
		Op:   op,
		Data: data,
	}
	if from != nil {
		msg.From = from.id
		msg.Name = from.name
	}
	u.msgs <- &msg
}

// leave cleans up after a user has closed his connection
func (u *User) leave() {
	u.room.Lock()
	delete(u.room.users, u)
	u.room.Unlock()
	u.room.mon <- &Message{Op: "leave"}
	u.room.send("leave", u.name, u)
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
	u.listVideos()
}

// listVideos sends a list of all available videos back to the client
// that requested it
func (u *User) listVideos() {
	u.send("list", videos, nil)
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
