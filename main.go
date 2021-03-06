//go:generate go-bindata -o static.go room.html

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// parse command line options
	addr := flag.String("addr", ":8080", "address to listen on")
	auth := flag.String("auth", "", "basic auth username and password (separated with \":\")")
	names := flag.String("words", "/usr/share/dict/words", "word-file to user for names")
	indexFile := flag.String("index", "", "file to serve as an index page")
	debug := flag.Bool("debug", false, "turn debugging mode on")
	dir := flag.String("dir", "", "directory to serve")
	flag.StringVar(&useWiki, "wiki", "", "use wiki prefix link for names")
	flag.Parse()

	// attempt to find index asset
	var err error
	index := []byte(`<!DOCTYPE html>
<title>Ijod!</title>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width" />
<pre><strong>IJOD:</strong> <a href="/new">new room</a>

(─‿‿─) .oO(<em>read the <a href="https://git.sr.ht/~zge/ijod">source</a>!</em>)`)
	if *indexFile != "" {
		index, err = ioutil.ReadFile(*indexFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// load valid authentications
	if *auth != "" {
		auths, err = loadAuthFile(*auth)
		if err != nil {
			log.Fatal(err)
		}
	}

	// change directory
	if *dir != "" {
		err := os.Chdir(*dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	// enable debugging mode, if requested
	if *debug {
		log.SetFlags(log.LUTC | log.Lshortfile | log.Ltime)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	// load file with words for words in it
	err = loadNames(*names)
	if err != nil {
		log.Printf("Words file not found (%s), defaulting to IDs", err)
	} else {
		log.Printf("Read in %d names from %s\n", len(words), *names)
	}

	// listen to signals to refresh video tree
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go refresher(c)

	// initialise and start HTTP server
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(".")) // TODO: prevent directory listing
	mux.Handle("/data/", http.StripPrefix("/data/", fs))
	mux.HandleFunc("/socket", socket)
	mux.HandleFunc("/room", room)
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, newRoom(), http.StatusFound)
	})

	var handler http.Handler
	if *auth != "" {
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "text/html")
				w.Write(index)
				return
			}

			if checkRequest(r) {
				mux.ServeHTTP(w, r)
			} else {
				w.Header().Set("WWW-Authenticate", `Basic realm="auth"`)
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			}
		})
	} else {
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.Header().Add("Content-Type", "text/html")
				w.Write(index)
				return
			}

			mux.ServeHTTP(w, r)
		})
	}
	log.Println("Listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, handler))
}
