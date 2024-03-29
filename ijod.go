package ijod

import (
	"embed"
	"net/http"

	"ijod/room"
	"ijod/tree"
	"ijod/user"
	"ijod/ytdl"
)

//go:embed ijod.js index.html style.css
var static embed.FS

// Merge all functionality into a HTTP handler
func Handler() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/room", room.Display)
	mux.HandleFunc("/socket", user.Socket)
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		url.Path = "/room"
		query := url.Query()
		query.Add("id", room.Create())
		url.RawQuery = query.Encode()
		http.Redirect(w, r, url.String(), http.StatusFound)
	})
	mux.Handle("/data/", http.StripPrefix("/data/", http.HandlerFunc(tree.Host)))
	if ytdl.Handler != nil {
		mux.Handle("/dl/", http.StripPrefix("/dl/", ytdl.Handler))
	}

	mux.Handle("/", http.FileServer(http.FS(static)))

	return mux
}
