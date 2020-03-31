package tigerserver

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err != nil {
		fmt.Fprint(c.out, "Please enter a number!")
		return
	}

	c.game.Start(numberOfPlayers, os.Stdout)

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
