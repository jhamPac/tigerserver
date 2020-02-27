package tigerserver

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore implements the PlayerStore interface for TigerServer
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

// NewFileSystemPlayerStore is a constructor for creating new FileSystemPlayerStore
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{database: database, league: league}
}

// GetLeague returns a slice of type Player
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

// GetPlayerScore takes a player's name and returns the score of the player specified
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin incremts the win data for specific player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
