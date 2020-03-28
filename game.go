package tigerserver

// Game any type that allows CLI to play a game
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
