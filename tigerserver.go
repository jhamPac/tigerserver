package tigerserver

import (
	"fmt"
	"net/http"
)

// TigerServer main server struct.
type TigerServer struct {
	Store  PlayerStore
	router *http.ServeMux
}

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// CreateTigerServer is the factory for the main server that creates and sets up routing too
func CreateTigerServer(store PlayerStore) *TigerServer {
	t := &TigerServer{
		store,
		http.NewServeMux(),
	}

	t.router.Handle("/", http.HandlerFunc(t.homeHandler))
	t.router.Handle("/league", http.HandlerFunc(t.leagueHandler))
	t.router.Handle("/players/", http.HandlerFunc(t.playersHandler))

	return t
}

func (t *TigerServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to Tiger Server!")
}

func (t *TigerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.router.ServeHTTP(w, r)
}

func (t *TigerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (t *TigerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
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
	player := r.URL.Path[len("/players/"):]
	return player
}
