package main

import (
	"fmt"
	"log"
	"net/http"

	ts "github.com/jhampac/tigerserver"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := ts.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := ts.NewTexas(ts.BlindAlerterFunc(ts.Alerter), store)

	server, err := ts.CreateTigerServer(store, game)
	if err != nil {
		fmt.Printf("TigerServer returned an error %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
