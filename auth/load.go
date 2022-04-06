package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Auth struct {
	hash string
}

type Keychain struct {
	data map[string]Auth
}

func Load(file string) (*Keychain, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	a := Keychain{make(map[string]Auth)}

	line := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		parts := strings.Split(s.Text(), "\t")
		if len(parts) != 2 {
			return nil, fmt.Errorf(
				"error: auth file too many columns on line %d",
				line)
		}

		a.data[parts[0]] = Auth{parts[1]}
		line++
	}

	return &a, s.Err()
}
