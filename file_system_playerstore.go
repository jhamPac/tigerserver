package tigerserver

import (
	"encoding/json"
	"io"
)

// FileSystemPlayerStore implements the PlayerStore interface for TigerServer
type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

// GetLeague returns a slice of type Player
func (f *FileSystemPlayerStore) GetLeague() League {
	f.Database.Seek(0, 0)
	league, _ := NewLeague(f.Database)
	return league
}

// GetPlayerScore takes a player's name and returns the score of the player specified
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

// RecordWin incremts the win data for specific player
func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.Database.Seek(0, 0)
	json.NewEncoder(f.Database).Encode(league)
}
