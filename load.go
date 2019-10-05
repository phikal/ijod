package main

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

const aclFile = ".ijod"

var (
	vlock  sync.Mutex
	videos *Directory
)

// Node is either a Video or a Directory
type Node interface {
	MarshalJSON() ([]byte, error)
}

func refresher(c <-chan os.Signal) {
	for {
		err := loadVideos()
		if err != nil {
			log.Println(err)
		}

		for _, r := range rooms {
			for u := range r.users {
				u.listVideos()
			}
		}
		<-c
	}
}

func walker(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	var node Node
	if info.IsDir() {
		node = NewDirectory(filepath.Join(p, aclFile))
	} else {
		node = (*Video)(&p)
	}
	videos.set(p, node)

	return nil
}

func loadVideos() error {
	vlock.Lock()
	defer vlock.Unlock()

	videos = NewDirectory("")
	return filepath.Walk(".", walker)
}
