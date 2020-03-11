package tigerserver

import "io"

// CLI can make calls to the server via terminal client
type CLI struct {
	store PlayerStore
	in    io.Reader
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	c.store.RecordWin("Cable")
}
