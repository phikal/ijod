//go:generate go-bindata -o static.go index.html

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
		pass  string
		debug bool
	)

	// configure flag parsing
	flag.BoolVar(&debug, "debug", false, "turn debugging mode on")
	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
	flag.StringVar(&pass, "pass", "", "basic auth password to require")
	flag.Parse()

	// enable debugging mode, if requested
	if debug {
		log.SetFlags(log.LUTC | log.Lshortfile | log.Ltime)
	} else {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	// prepare index.html
	index, err := Asset("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// initialize and start HTTP server
	fs := http.FileServer(http.Dir("."))
	http.Handle("/data/", http.StripPrefix("/data/", fs))
	http.HandleFunc("/socket", socket)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "no such site", http.StatusNotImplemented)
			return
		}

		if _, err := w.Write(index); err != nil {
			log.Println(err)
		}
	})

	var handler http.Handler
	if pass != "" {
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, try, _ := r.BasicAuth()
			if try != pass {
				w.Header().Set("WWW-Authenticate", `Basic realm="pass"`)
				http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			} else {
				http.DefaultServeMux.ServeHTTP(w, r)
			}
		})
	}
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
