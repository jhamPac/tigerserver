package tigerserver

import "io"

// Game any type that allows CLI to play a game
type Game interface {
	Start(numberOfPlayers int, alerterDestination io.Writer)
	Finish(winner string)
}
