package main

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

var (
	cvid    *Video
	vlock   sync.Mutex
	waiting = make(chan *User)
)

type Video struct {
	sync.Mutex
	path    string
	playing bool
	updated time.Time
}

func (v *Video) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.path)
}

func selectVideo(name string, from *User) error {
	vlock.Lock()
	defer vlock.Unlock()

	var ok bool
	if cvid, ok = fpVideos[name]; !ok {
		return errors.New("No such video: " + name)
	} else {
		if ok {
			send("select", name, from)
		}
	}

	return nil
}

// play stops the current video, at the current position, and doesn't
// change anything if it isn't playing
//
// a "pause"-signal is sent to all members of this room
func (v *Video) pause(from *User) {
	if v == nil || !v.playing {
		return
	}

	send("pause", nil, from)
	v.playing = false
}

// play continues the video, if it is been paused, and doesn't change
// anything if it's playing.
//
// a "play"-signal is sent to all members of this room
func (v *Video) play(from *User) {
	if v == nil || v.playing {
		return
	}

	send("play", nil, from)
	v.playing = true
}

// jumpTo sets a absolute position in the video
func (v *Video) jumpTo(pos float64, from *User) {
	if v == nil {
		return
	}

	v.Lock()
	defer v.Unlock()

	if time.Since(v.updated) < time.Millisecond*10 {
		return
	}

	v.pause(from)
	v.updated = time.Now()
	send("time", pos, from)
}

func (v *Video) startp() {
	v.Lock()
	defer v.Unlock()

	all := true
	for u := range users {
		all = all && u.ready
	}

	if all {
		for u := range users {
			u.ready = false
		}
		v.play(nil)
	}
}
