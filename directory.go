package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
)

type Directory struct {
	acl  map[string]bool
	tree map[string]Node
}

// NewDirectory creates a directory object
func NewDirectory(aclfile string) *Directory {
	return &Directory{
		loadACL(aclfile),
		make(map[string]Node),
	}
}

func (d *Directory) MarshalJSON() ([]byte, error) {
	if d == nil {
		return nil, nil
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(d.tree)
	return buf.Bytes(), nil
}

// Allows knows if a user is permitted to see this directory or not
func (d *Directory) allows(u *User) bool {
	return d.acl == nil || d.acl[u.auth]
}

func topDir(path string) (string, string) {
	parts := strings.SplitN(path, string(filepath.Separator), 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func (d *Directory) find(path string) *Video {
	next, rest := topDir(path)
	node, ok := d.tree[next]
	if ok {
		switch node.(type) {
		case *Video:
			return node.(*Video)
		case *Directory:
			dir := node.(*Directory)
			return dir.find(rest)
		}
	}
	return nil
}

func (d *Directory) set(path string, nd Node) {
	next, rest := topDir(path)
	if next == "." {
		return
	}
	node := d.tree[next]
	dir, ok := node.(*Directory)
	if !ok {
		if rest == "" {
			d.tree[next] = nd
			return
		}

		dir = NewDirectory(filepath.Join(path, aclFile))
		d.tree[next] = dir
	}
	dir.set(rest, nd)

}

func (d *Directory) filterFor(u *User) interface{} {
	fd := make(map[string]interface{})
	for path, node := range d.tree {
		switch node.(type) {
		case *Video: // add all videos from a legal tree
			if filepath.Base(path) != aclFile {
				fd[path] = node
			}
		case *Directory: // traverse subtrees recursivly
			sd := node.(*Directory)
			if sd.allows(u) {
				fd[path] = sd.filterFor(u)
			}
		}
	}
	return fd
}
