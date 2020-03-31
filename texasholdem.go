package tigerserver

import (
	"io"
	"time"
)

// TexasHoldem defines the function and boundry of the game being played
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewTexas creates a card game that implements the Game interface
func NewTexas(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{alerter: alerter, store: store}
}

// Start invokes a game and blindalerter call
func (t *TexasHoldem) Start(numberOfPlayers int, to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		t.alerter.ScheduleAlertAt(blindTime, blind, to)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends a game with a winner
func (t *TexasHoldem) Finish(winner string) {
	t.store.RecordWin(winner)
}
