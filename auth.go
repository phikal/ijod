package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Auth represents a valid login
type Auth struct {
	user string
	pass string
}

var auths []*Auth

func loadAuthFile(file string) (a []*Auth, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	s := bufio.NewScanner(f)
	line := 1
	for s.Scan() {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) != 2 {
			err = fmt.Errorf("error: auth file too many columns on line %d", line)
			return
		}

		auth := &Auth{user: parts[0], pass: parts[1]}
		a = append(a, auth)
		line++
	}
	err = s.Err()
	return
}

func checkAuth(a *Auth, user, pass string) bool {
	sha := sha256.Sum256([]byte(user + pass))
	sum := hex.EncodeToString(sha[:])
	return user == a.user && sum == a.pass
}

func checkAnyAuth(user, pass string) *Auth {
	for _, a := range auths {
		if checkAuth(a, user, pass) {
			return a
		}
	}
	return nil
}

func checkRequest(r *http.Request) bool {
	user, pass, _ := r.BasicAuth()
	return checkAnyAuth(user, pass) != nil
}

func loadACL(path string) map[string]bool {
	if path == "" {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
		}
		return nil
	}
	defer file.Close()
	acl, err := parseACL(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	return acl
}

func parseACL(in io.Reader) (map[string]bool, error) {
	var (
		names = make(map[string]bool)
		buf   = bufio.NewScanner(in)
	)
	for buf.Scan() {
		names[buf.Text()] = true
	}
	return names, buf.Err()
}
