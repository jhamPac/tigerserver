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

	fmt.Println("Let's enter the MATRIX")
	fmt.Println("Type {Name} wins to record a win")

	game := ts.NewCLI(store, os.Stdin, ts.BlindAlerterFunc(ts.StdOutAlerter))
	game.PlayPoker()
}
