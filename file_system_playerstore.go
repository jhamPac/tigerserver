package tigerserver

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

// FileSystemPlayerStore implements the PlayerStore interface for TigerServer
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

// FileSystemPlayerStoreFromFile takes a path and returns a fully configured FileSystemPlayerStore
func FileSystemPlayerStoreFromFile(path string) (*FileSystemPlayerStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store, err := NewFileSystemPlayerStore(db)

	if err != nil {
		return nil, nil, fmt.Errorf("problem creating file system player store, %v", err)
	}

	return store, closeFunc, nil
}

// NewFileSystemPlayerStore is a constructor for creating new FileSystemPlayerStore
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{database: json.NewEncoder(&tape{file}), league: league}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)
	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}

	return nil
}

// GetLeague returns a slice of type Player
func (f *FileSystemPlayerStore) GetLeague() League {
	// sort by most wins
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

// GetPlayerScore takes a player's name and returns the score of the player specified
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	n := strings.ToLower(name)
	player := f.league.Find(n)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin incremts the win data for specific player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	n := strings.ToLower(name)
	player := f.league.Find(n)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{Name: n, Wins: 1})
	}
	f.database.Encode(f.league)
}
