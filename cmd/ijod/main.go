package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"ijod"
	"ijod/auth"
	"ijod/user"
)

var (
	listen = flag.String("listen", ":8080", "address to listen on")
	authd  = flag.String("auth", "", "basic auth username and password (separated with \":\")")
	names  = flag.String("names", "/usr/share/dict/words", "word-file to user for names")
	dir    = flag.String("dir", ".", "directory to serve")
	debug  = flag.Bool("debug", false, "turn debugging mode on")
)

func main() {
	var err error

	flag.Parse()

	// Load file with words for words in it
	if *names != "" {
		err = user.LoadNames(*names)
		if err != nil {
			log.Fatalf("Failed to read %s: %s", *names, err)
		}
	}

	// Load authentication data
	var kc *auth.Keychain
	if *authd != "" {
		var err error
		kc, err = auth.Load(*authd)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Change directory before starting the server
	err = os.Chdir(*dir)
	if err != nil {
		log.Fatal(err)
	}

	// Enable debugging mode, if requested
	if *debug {
		log.SetFlags(log.Llongfile)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	// Start the server
	log.Printf("Listening on http://localhost%s", *listen)
	log.Fatal(http.ListenAndServe(*listen, kc.Wrap(ijod.Handler())))
}
