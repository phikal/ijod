package tree

import (
	"net/http"
)

func Host(w http.ResponseWriter, r *http.Request) {
	// TODO: Disable directory listing
	http.FileServer(http.Dir(".")).ServeHTTP(w, r)
}
