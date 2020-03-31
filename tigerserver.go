package tigerserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

// TigerServer main server struct.
type TigerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
}

// PlayerStore for server methods
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() League
}

// Player represents a user entity
type Player struct {
	Name string
	Wins int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const htmlTemplatePath = "./game.html"

// CreateTigerServer is the factory for the main server that creates and sets up routing too
func CreateTigerServer(store PlayerStore) (*TigerServer, error) {
	t := new(TigerServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	t.template = tmpl
	t.store = store

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(t.homeHandler))
	router.Handle("/league", http.HandlerFunc(t.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(t.playersHandler))
	router.Handle("/game", http.HandlerFunc(t.game))
	router.Handle("/ws", http.HandlerFunc(t.webSocket))

	t.Handler = router

	return t, nil
}

func (t *TigerServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to Tiger Server!")
}

func (t *TigerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(t.store.GetLeague())
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

func (t *TigerServer) game(w http.ResponseWriter, r *http.Request) {
	t.template.Execute(w, nil)
}

func (t *TigerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	t.store.RecordWin(string(winnerMsg))
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
