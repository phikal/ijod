package room

import (
	"log"

	"ijod/mesg"
)

func (r *Room) Join(id *mesg.User) {
	log.Println(id.Name, "to join", r.Name)
	r.enter <- id
}

func (r *Room) Leave(id *mesg.User) {
	log.Println(id.Name, "to leave", r.Name)
	r.leave <- id
}
