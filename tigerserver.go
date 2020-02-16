package tigerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// TigerServer main server struct.
type TigerServer struct {
	store PlayerStore
}

func (t *TigerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, t.store.GetPlayerScore(player))
}
