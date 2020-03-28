package tigerserver

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt used to ask the user a question on CLI start up
const PlayerPrompt = "Please enter the number of player: "

// CLI can make calls to the server via terminal client
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI factory function for object
func NewCLI(i io.Reader, o io.Writer, g Game) *CLI {
	return &CLI{in: bufio.NewScanner(i), out: o, game: g}
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayersInput := c.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner := extractWinner(winnerInput)

	c.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
