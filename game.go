package tigerserver

import "time"

// Game defines the function and boundry of the game being played
type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

// Start invokes a game and blindalerter call
func (g *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends a game with a winner
func (g *Game) Finish(winner string) {
	g.store.RecordWin(winner)
}
