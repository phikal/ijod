package auth

import (
	"net/http"
)

func (k *Keychain) Wrap(handler http.Handler) http.Handler {
	if k == nil {
		return handler
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if k.CheckRequest(r) {
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="auth"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		}
	})
}
