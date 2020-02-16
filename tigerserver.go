package tigerserver

import (
	"fmt"
	"net/http"
)

// Roar response handler for HTTP.
func Roar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
