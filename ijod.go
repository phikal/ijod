package ijod

import (
	"embed"
	"net/http"

	"ijod/room"
	"ijod/tree"
	"ijod/user"
)

//go:embed ijod.js index.html style.css
var static embed.FS

// Merge all functionality into a HTTP handler
func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/room", room.Display)
	mux.HandleFunc("/socket", user.Socket)
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/room?id="+room.Create(), http.StatusFound)
	})
	mux.Handle("/data/", http.StripPrefix("/data/", http.HandlerFunc(tree.Host)))
	mux.Handle("/", http.FileServer(http.FS(static)))

	return mux
}
