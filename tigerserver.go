package tigerserver

import (
	"fmt"
	"net/http"
)

// TigerServer initiates an instance of a power server.
func TigerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
