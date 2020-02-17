package tigerserver

import (
	"fmt"
	"net/http"
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
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			t.processWin(w, r)
		case http.MethodGet:
			t.showScore(w, r)
		}
	}))

	router.ServeHTTP(w, r)
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
	player := r.URL.Path[len("/players/"):]
	return player
}
