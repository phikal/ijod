package tree

import (
	"io/fs"
	"path/filepath"

	"ijod/mesg"
)

// Generate a "files" message with a tree
func Message() *mesg.Message {
	type node map[string]interface{}

	flat := node{".": make(node)}
	filepath.WalkDir(".", fs.WalkDirFunc(
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

	// log.Print(flat)
	return &mesg.Message{
		Type: "files",
		Data: flat["."],
	}
}
