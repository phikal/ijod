package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

// Video stores the information about one video in one room
type Video string

func (v *Video) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(string(*v))
	return buf.Bytes(), nil
}

func (r *Room) selectVideo(name string, from *User) error {
	if r.hasAdmin && r.admin != from {
		return nil
	}

	r.Lock()
	defer r.Unlock()

	vid := videos.find(name)
	if vid == nil {
		return errors.New("No such video: " + name)
	}
	r.send("select", name, from)
	return nil
}

// play stops the current video, at the current position, and doesn't
// change anything if it isn't playing
//
// a "pause"-signal is sent to all members of this room
func (r *Room) pause(from *User) {
	if r.hasAdmin && r.admin != from {
		return
	}

	if r.vid == nil || !r.playing {
		return
	}

	r.send("pause", nil, from)
	r.playing = false
}

// play continues the video, if it is been paused, and doesn't change
// anything if it's playing.
//
// a "play"-signal is sent to all members of this room
func (r *Room) play(from *User) {
	if r.hasAdmin && r.admin != from {
		return
	}

	if r.vid == nil || r.playing {
		return
	}

	r.send("play", nil, from)
	r.playing = true
}

// jumpTo sets a absolute position in the video
func (r *Room) jumpTo(pos float64, from *User) {
	if r.hasAdmin && r.admin != from {
		return
	}

	if r.vid == nil {
		return
	}

	if time.Since(r.updated) < time.Millisecond*500 {
		return
	}

	r.updated = time.Now()
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
