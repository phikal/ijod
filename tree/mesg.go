package tree

import (
	"io/fs"
	"log"
	"path/filepath"

	"ijod/mesg"
)

// Generate a "files" message with a tree
func Message() *mesg.Message {
	type (
		file struct {
			Name string `json:"name"`
			Dura string `json:"duration"`
		}
		node map[string]interface{}
	)

	flat := node{".": make(node)}
	err := filepath.WalkDir(".", fs.WalkDirFunc(
		func(path string, d fs.DirEntry, err error) error {
			var ent interface{}

			if d.IsDir() {
				ent = make(node)
			} else {
				ent = path
			}

			sup := filepath.Dir(path)
			if parent, ok := flat[sup].(node); ok {
				parent[filepath.Base(path)] = ent
			}
			flat[path] = ent

			return nil
		}))
	if err != nil {
		log.Print(err)
		return nil
	}

	return &mesg.Message{
		Type: "files",
		Data: flat["."],
	}
}
