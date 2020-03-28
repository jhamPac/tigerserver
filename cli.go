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
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

// NewCLI factory function for object
func NewCLI(store PlayerStore, i io.Reader, o io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store: store, in: bufio.NewScanner(i), out: o, alerter: alerter}
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)
	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)

	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
