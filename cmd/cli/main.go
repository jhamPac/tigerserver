package main

import (
	"fmt"
	"log"
	"os"

	ts "github.com/jhampac/tigerserver"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := ts.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := ts.NewTexas(ts.BlindAlerterFunc(ts.StdOutAlerter), store)
	cli := ts.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play a game")
	fmt.Println("Type any name to record a win!")
	cli.PlayPoker()
}
