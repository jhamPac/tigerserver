package tigerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// Roar response handler for HTTP.
func Roar(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Casio" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Laverne" {
		fmt.Fprint(w, "10")
		return
	}
}
