package web

import (
	"fmt"
	"html"
	"net/http"
)

type Recommended struct {
}

func (rec Recommended) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rec.handleGet(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// TODO: update my tiles
// TODO: return json
func (rec Recommended) handleGet(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
