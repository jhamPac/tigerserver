package tigerserver

// NewInMemoryPlayerStore is a factory that returns InMemoryPlayerStore
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
