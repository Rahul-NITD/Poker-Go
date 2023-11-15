package main

import (
	"log"
	"net/http"

	poker "github.com/Rahul-NITD/Poker"
)

func main() {
	storage, close, _ := poker.NewDBStorage(false)
	defer close(storage)
	server := poker.NewServer(storage)
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(server.ServeHTTP)))
}
