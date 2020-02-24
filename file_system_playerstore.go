package tigerserver

import (
	"io"
)

// FileSystemPlayerStore implements the PlayerStore interface for TigerServer
type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

// GetLeague returns a slice of type Player
func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

// GetPlayerScore takes a player's name and returns the score of the player specified
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}

	return wins
}
