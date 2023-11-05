package room

import (
	"log"
	"os"
	"sync"

	"ijod/mesg"
)

var rooms sync.Map

type Room struct {
	Name  string
	enter chan *mesg.User
	leave chan *mesg.User
	files []string
}

func Create() string {
	room := Room{
		Name:  randName(),
		enter: make(chan *mesg.User),
		leave: make(chan *mesg.User),
	}
	rooms.Store(room.Name, &room)
	go room.daemon()

	return room.Name
}

func (r *Room) Forget() {
	for _, file := range r.files {
		os.Remove(file)
	}

	rooms.Delete(r.Name)
	log.Println("Forget room", r.Name)
}

func GetRoom(id string) (*Room, bool) {
	room, ok := rooms.Load(id)
	return room.(*Room), ok
}
