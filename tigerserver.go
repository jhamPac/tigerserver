package tigerserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TigerServer main server struct.
type TigerServer struct {
	store PlayerStore
	http.Handler
}

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// Player represents a user entity
type Player struct {
	Name string
	Wins int
}

// CreateTigerServer is the factory for the main server that creates and sets up routing too
func CreateTigerServer(store PlayerStore) *TigerServer {
	t := new(TigerServer)

	t.store = store

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(t.homeHandler))
	router.Handle("/league", http.HandlerFunc(t.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(t.playersHandler))

	t.Handler = router

	return t
}

func (t *TigerServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to Tiger Server!")
}

func (t *TigerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(t.getLeagueTable())
	w.WriteHeader(http.StatusOK)
}

func (t *TigerServer) getLeagueTable() []Player {
	return []Player{
		{"Myer", 10},
		{"Baker", 2},
	}
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
	t.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (t *TigerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := trimPlayerURL(r)
	score := t.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func trimPlayerURL(r *http.Request) string {
	player := r.URL.Path[len("/players/"):]
	return player
}
