package tigerserver

import (
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanData := createTempFile(t, `[
				{"Name": "storm", "Wins": 10},
				{"Name": "rogue", "Wins": 30},
				{"Name": "xavior", "Wins": 50}
			]`)

	defer cleanData()

	store, err := NewFileSystemPlayerStore(database)

	if err != nil {
		t.Fatalf("store construction failed with: %v", err)
	}

	t.Run("/league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"xavior", 50},
			{"rogue", 30},
			{"storm", 10},
		}

		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get a player score from a league", func(t *testing.T) {
		got := store.GetPlayerScore("Rogue")
		want := 30

		if got != want {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Storm")
		got := store.GetPlayerScore("Storm")
		want := 11

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Cyclops")
		got := store.GetPlayerScore("Cyclops")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("league sorted", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"xavior", 50},
			{"rogue", 30},
			{"storm", 11},
			{"cyclops", 1},
		}

		assertLeague(t, got, want)
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}
