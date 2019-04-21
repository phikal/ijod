package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var (
	videos  map[string]*Video
	cvid    *Video
	vlock   sync.Mutex
	waiting = make(chan *User)
)

type Video struct {
	sync.Mutex
	path     string
	playing  bool
	fsyncing bool
	updated  time.Time
}

func init() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGUSR1)

		for {
			<-c
			err := loadVideos()
			if err != nil {
				log.Fatal(err)
			}

			for u := range users {
				u.listVideos()
			}
		}
	}()

	loadVideos()
}

func loadVideos() error {
	vlock.Lock()
	defer vlock.Unlock()

	videos = make(map[string]*Video)
	return filepath.Walk(".", func(path string,
		info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			videos[path] = &Video{path: path}
		}
		return err
	})
}

func selectVideo(name string, from *User) error {
	vlock.Lock()
	defer vlock.Unlock()

	if v, ok := videos[name]; !ok {
		return errors.New("No such video: " + name)
	} else {
		if cvid != nil {
			cvid.fsyncing = false
		}
		cvid = v
		send("select", name, from)
	}

	return nil
}

// play stops the current video, at the current position, and doesn't
// change anything if it isn't playing
//
// a "pause"-signal is sent to all members of this room
func (v *Video) pause(from *User) {
	if v == nil || !v.playing || v.fsyncing {
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
	if v == nil || v.playing || v.fsyncing {
		return
	}

	send("play", nil, from)
	v.playing = true
}

// jumpTo sets a absolute position in the video
func (v *Video) jumpTo(pos float64, from *User) {
	if v == nil || v.fsyncing {
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
		v.fsyncing = false
		v.play(nil)
	}
}
