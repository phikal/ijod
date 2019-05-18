package main

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

// Video stores the information about one video in one room
type Video struct {
	sync.Mutex
	path    string
	playing bool
	updated time.Time
}

// MarshalJSON transforms a Video into a JSON object, by turning it into
// it's path-value as a string
func (v *Video) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.path)
}

func (r *Room) selectVideo(name string, from *User) error {
	r.Lock()
	defer r.Unlock()

	var ok bool
	if r.vid, ok = fpVideos[name]; !ok {
		return errors.New("No such video: " + name)
	} else if ok {
		r.send("select", name, from)
	}

	return nil
}

// play stops the current video, at the current position, and doesn't
// change anything if it isn't playing
//
// a "pause"-signal is sent to all members of this room
func (r *Room) pause(from *User) {
	if r.vid == nil || !r.vid.playing {
		return
	}

	r.send("pause", nil, from)
	r.vid.playing = false
}

// play continues the video, if it is been paused, and doesn't change
// anything if it's playing.
//
// a "play"-signal is sent to all members of this room
func (r *Room) play(from *User) {
	if r.vid == nil || r.vid.playing {
		return
	}

	r.send("play", nil, from)
	r.vid.playing = true
}

// jumpTo sets a absolute position in the video
func (r *Room) jumpTo(pos float64, from *User) {
	if r.vid == nil {
		return
	}

	r.vid.Lock()
	defer r.vid.Unlock()

	if time.Since(r.vid.updated) < time.Millisecond*500 {
		return
	}

	r.vid.updated = time.Now()
	r.pause(from)
	r.send("time", pos, from)
}

func (r *Room) startp() {
	r.Lock()
	defer r.Unlock()

	all := true
	for u := range r.users {
		all = all && u.ready
	}

	if all {
		for u := range r.users {
			u.ready = false
		}
		r.play(nil)
	}
}
