package room

import (
	"log"
	"os"
	"sync"

	"ijod/mesg"
)

var (
	lock  sync.Mutex
	rooms = make(map[string]*Room)
)

type Room struct {
	Name  string
	enter chan *mesg.User
	leave chan *mesg.User
	files []string
}

func Create() string {
	defer lock.Unlock()
	lock.Lock()

	room := Room{
		Name:  randName(),
		enter: make(chan *mesg.User),
		leave: make(chan *mesg.User),
	}
	rooms[room.Name] = &room
	go room.daemon()

	return room.Name
}

func (r *Room) Forget() {
	defer lock.Unlock()
	lock.Lock()

	for _, file := range r.files {
		os.Remove(file)
	}

	delete(rooms, r.Name)
	log.Println("Forget room", r.Name)
}

func GetRoom(id string) (*Room, bool) {
	defer lock.Unlock()
	lock.Lock()

	room, ok := rooms[id]
	return room, ok
}
