package main

import (
	"fmt"
	"log"
	"os"

	ts "github.com/jhampac/tigerserver"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's enter the MATRIX")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := ts.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}

	game := ts.CLI{store, os.Stdin}
	game.PlayPoker()
}
