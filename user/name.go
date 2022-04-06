package user

import (
	"bufio"
	"compress/gzip"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

var (
	counter int64
	words   []string
)

func LoadNames(name string) (err error) {
	var file io.ReadCloser
	file, err = os.Open(name)
	if err != nil {
		return
	}

	// Check if file is compressed
	if strings.HasSuffix(name, ".gz") {
		file, err = gzip.NewReader(file)
		if err != nil {
			return
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		words = append(words, name)
	}

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return scanner.Err()
}

func nextName() string {
	var word string

	if words != nil {
		word = words[int(counter)%len(words)]
	} else {
		word = strconv.FormatInt(counter, 10)
	}
	atomic.AddInt64(&counter, counter+rand.Int63n(8))

	return word
}
