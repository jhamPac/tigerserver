package tigerserver

import (
	"bufio"
	"io"
	"strings"
)

// CLI can make calls to the server via terminal client
type CLI struct {
	store PlayerStore
	in    io.Reader
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	reader := bufio.NewScanner(c.in)
	reader.Scan()
	c.store.RecordWin(extractWinner(reader.Text()))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
