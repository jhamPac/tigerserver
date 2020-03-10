package tigerserver

// CLI can make calls to the server via terminal client
type CLI struct {
	store PlayerStore
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	c.store.RecordWin("Cable")
}
