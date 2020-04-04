package tigerserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var dummyGame = &GameSpy{}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Casio":   20,
			"Laverne": 10,
		},
	}

	server, _ := CreateTigerServer(&store, dummyGame)

	t.Run("returns Casio's score", func(t *testing.T) {
		request := newGetScoreRequest("Casio")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns Laverne's score", func(t *testing.T) {
		request := newGetScoreRequest("Laverne")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Jupitar")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got status %d but wanted %d", got, want)
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{map[string]int{}, []string{}, nil}

	server, _ := CreateTigerServer(&store, dummyGame)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Casio"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		assertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Wolverine", 22},
			{"Cyclops", 7},
			{"Beast", 10},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server, _ := CreateTigerServer(&store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, "application/json")

	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		store := &StubPlayerStore{}
		server, _ := CreateTigerServer(store, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("start game with 3 players and declare Beast the winner", func(t *testing.T) {
		game := &GameSpy{}
		store := &StubPlayerStore{}
		winner := "Beast"
		ts, _ := CreateTigerServer(store, game)
		server := httptest.NewServer(ts)
		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})
}
