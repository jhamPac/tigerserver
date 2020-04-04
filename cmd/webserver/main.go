package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jhampac/tigerserver"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := tigerserver.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := tigerserver.NewTexas(tigerserver.BlindAlerterFunc(tigerserver.Alerter), store)

	server, err := tigerserver.New(store, game)
	if err != nil {
		fmt.Printf("TigerServer returned an error %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
