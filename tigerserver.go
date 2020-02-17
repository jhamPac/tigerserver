package tigerserver

import (
	"fmt"
	"net/http"
	"strings"
)

// TigerServer main server struct.
type TigerServer struct {
	Store PlayerStore
}

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

func (t *TigerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		t.processWin(w, r)
	case http.MethodGet:
		t.showScore(w, r)
	}
}

func (t *TigerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := trimPlayerURL(r)
	t.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (t *TigerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := trimPlayerURL(r)
	score := t.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func trimPlayerURL(r *http.Request) string {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	return player
}
