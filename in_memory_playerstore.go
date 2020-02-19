package tigerserver

// NewInMemoryPlayerStore is a factory that returns InMemoryPlayerStore; great example of a factory
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{make(map[string]int)}
}

// InMemoryPlayerStore for a DB agnostic store
type InMemoryPlayerStore struct {
	Store map[string]int
}

// RecordWin records the win for a specified player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.Store[name]++
}

// GetPlayerScore returns the score for a specified player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.Store[name]
}

// GetLeague returns a slice of Player of in the leagues
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.Store {
		league = append(league, Player{name, wins})
	}

	return league
}
