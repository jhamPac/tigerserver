package tigerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// Roar response handler for HTTP.
func Roar(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(name string) string {
	if name == "Casio" {
		return "20"
	}

	if name == "Laverne" {
		return "10"
	}

	return ""
}
