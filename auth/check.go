package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

// Check if USER may pass with PASS
func (k *Keychain) Check(user, pass string) bool {
	// Query keychain
	auth, ok := k.data[user]
	if !ok {
		return false
	}

	// Calculate a hash
	sha := sha256.Sum256([]byte(user + pass)) // TODO: use bcrypt
	sum := hex.EncodeToString(sha[:])

	// Compare the hash values
	return sum == auth.hash
}

// Check the basic auth of a HTTP request
func (k *Keychain) CheckRequest(r *http.Request) bool {
	if k == nil || r.URL.Path == "/" {
		return true
	}

	user, pass, _ := r.BasicAuth()
	return k.Check(user, pass)
}
