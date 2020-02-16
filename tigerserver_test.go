package tigerserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Casio's score", func(t *testing.T) {
		request := newGetScoreRequest("Casio")
		response := httptest.NewRecorder()

		Roar(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns Laverne's score", func(t *testing.T) {
		request := newGetScoreRequest("Laverne")
		response := httptest.NewRecorder()

		Roar(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q wanted %q", got, want)
	}
}
