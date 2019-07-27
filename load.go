package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"
)

var (
	vlock  sync.Mutex
	videos map[string]interface{}
)

func init() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGUSR1)

		for {
			err := loadVideos()
			if err != nil {
				log.Println(err)
			}

			for _, r := range rooms {
				r.refreshVideos()
				for u := range r.users {
					u.listVideos()
				}
			}
			<-c
		}
	}()
}

func walkDir(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	var data interface{}
	if info.IsDir() {
		data = make(map[string]interface{})
	} else {
		data = p
	}

	var list []string
	for p != "." {
		list = append(list, path.Base(p))
		p = path.Dir(p)
	}
	list = append(list, p)

	var ok bool
	place := videos

	for i := len(list) - 1; i > 0; i-- {
		place, ok = place[list[i]].(map[string]interface{})
		if !ok {
			return errors.New("not a place")
		}
	}
	place[list[0]] = data

	return nil
}

func loadVideos() error {
	vlock.Lock()
	defer vlock.Unlock()

	videos = make(map[string]interface{})
	return filepath.Walk(".", walkDir)
}
