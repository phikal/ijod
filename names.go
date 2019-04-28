package main

import (
	"bufio"
	"math/rand"
	"os"
)

var words []string

func loadNames(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return err
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return scanner.Err()
}
