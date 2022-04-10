package user

import (
	"bufio"
	"compress/bzip2"
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
	defer file.Close()

	var r io.Reader
	// Check if file is compressed
	switch {
	case strings.HasSuffix(name, ".gz"):
		r, err = gzip.NewReader(file)
	case strings.HasSuffix(name, ".bz2"):
		r = bzip2.NewReader(file)
	default:
		r = file
	}
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(r)
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
	atomic.AddInt64(&counter, counter+rand.Int63n(8)+1)

	return word
}
