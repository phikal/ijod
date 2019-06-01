//go:generate go-bindata -o static.go room.html

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// parse command line options
	addr := flag.String("addr", ":8080", "address to listen on")
	auth := flag.String("auth", "", "basic auth username and password (separated with \":\")")
	names := flag.String("words", "/usr/share/dict/words", "word-file to user for names")
	indexFile := flag.String("index", "", "file to serve as an index page")
	debug := flag.Bool("debug", false, "turn debugging mode on")
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
		log.Fatalln(err)
	}
	log.Printf("Read in %d names\n", len(words))

	// initialise and start HTTP server
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
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

			tryUser, tryPass, _ := r.BasicAuth()
			if tryUser+":"+tryPass != *auth {
				w.Header().Set("WWW-Authenticate", `Basic realm="auth"`)
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			} else {
				mux.ServeHTTP(w, r)
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
