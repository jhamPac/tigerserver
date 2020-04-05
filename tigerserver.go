package tigerserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"

	"github.com/gorilla/websocket"
)

// TigerServer main server struct.
type TigerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
	game     Game
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

type playerServerWS struct {
	*websocket.Conn
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("error reading from websocket %v\n", err)
	}
	return string(msg)
}

func (w *playerServerWS) Write(p []byte) (n int, err error) {
	if err := w.WriteMessage(1, p); err != nil {
		return 0, err
	}

	return len(p), nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// path to the root of this project
var (
	_, b, _, _       = runtime.Caller(0)
	basepath         = filepath.Dir(b)
	htmlTemplatePath = basepath + "/game.html"
)

// New creates an initialized TigerServer
func New(store PlayerStore, game Game) (*TigerServer, error) {
	t := new(TigerServer)

	tmpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	t.template = tmpl
	t.store = store
	t.game = game

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(t.homeHandler))
	router.Handle("/league", http.HandlerFunc(t.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(t.playersHandler))
	router.Handle("/game", http.HandlerFunc(t.handleGame))
	router.Handle("/ws", http.HandlerFunc(t.webSocket))

	t.Handler = router

	return t, nil
}

func (t *TigerServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Welcome to =: Tiger Server :=")
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

func (t *TigerServer) handleGame(w http.ResponseWriter, r *http.Request) {
	t.template.Execute(w, nil)
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("problem upgrading connection to WebSockets %v\n", err)
	}
	return &playerServerWS{conn}
}

func (t *TigerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	nMsg := ws.WaitForMsg()
	nPlayers, _ := strconv.Atoi(nMsg)
	t.game.Start(nPlayers, ws)

	winner := ws.WaitForMsg()
	t.game.Finish(winner)
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
