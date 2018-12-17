package handlers

import (
	"net/http"
)

// Map function receive a context from incoming request and redirect to URL specify.
func Map(path map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path, exists := path[r.URL.Path]

		if exists {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}
