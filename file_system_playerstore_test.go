package tigerserver

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database, cleanData := createTempFile(t, `[
		{"Name": "Storm", "Wins": 10},
		{"Name": "Rogue", "Wins": 30
		}]`)

	defer cleanData()

	store := FileSystemPlayerStore{database}

	t.Run("/league from a reader", func(t *testing.T) {
		got := store.GetLeague()
		want := []Player{
			{"Storm", 10},
			{"Rogue", 30},
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
}

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
