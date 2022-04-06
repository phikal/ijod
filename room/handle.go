package room

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

func exists(id string) bool {
	lock.Lock()
	_, ok := rooms[id]
	lock.Unlock()
	return ok
}

//go:embed room.html
var room []byte

func Display(w http.ResponseWriter, r *http.Request) {
	if !exists(r.URL.Query().Get("id")) {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, `<!DOCTYPE html>
<title>Ijod?</title>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width" />
<pre>This room has been destroyed. Create a <a href="/new">new one</a>.`)
		return
	}

	_, err := w.Write(room)
	if err != nil {
		log.Println(err)
	}
}
