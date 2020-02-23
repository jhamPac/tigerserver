package tigerserver

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
			{"Name": "Storm", "Wins": 10},
			{"Name": "Rogue", "Wins": 30
			}]`)

	store := FileSystemPlayerStore{database}

	t.Run("/league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"Storm", 10},
			{"Rogue", 30},
		}

		assertLeague(t, got, want)
	})

	t.Run("get a player score from a league", func(t *testing.T) {
		got := store.GetPlayerScore("Rogue")
		want := 30

		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})
}
