package tigerserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Casio's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Casio/", nil)
		response := httptest.NewRecorder()

		Roar(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
