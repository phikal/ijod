//go:generate go-bindata -o static.go index.html room.html

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var (
		addr  string
		auth  string
		debug bool
	)

	// configure flag parsing
	flag.BoolVar(&debug, "debug", false, "turn debugging mode on")
	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
	flag.StringVar(&auth, "auth", "", "basic auth username and password (separated with \":\")")
	flag.Parse()

	// enable debugging mode, if requested
	if debug {
		log.SetFlags(log.LUTC | log.Lshortfile | log.Ltime)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	// attempt to find index asset
	index, err := Asset("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// initialise and start HTTP server
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("."))
	mux.Handle("/data/", http.StripPrefix("/data/", fs))
	mux.HandleFunc("/socket", socket)
	mux.HandleFunc("/room", room)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/new":
			http.Redirect(w, r, newRoom(), http.StatusFound)
		case "/":
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		default:
			http.Error(w, "no such site", http.StatusNotImplemented)
		}
	})

	var handler http.Handler = mux
	if auth != "" {
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			try_user, try_pass, _ := r.BasicAuth()
			if try_user+":"+try_pass != auth {
				w.Header().Set("WWW-Authenticate", `Basic realm="auth"`)
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			} else {
				mux.ServeHTTP(w, r)
			}
		})
	}
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
