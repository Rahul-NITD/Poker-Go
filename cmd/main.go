package main

import (
	"log"
	"net/http"

	poker "github.com/Rahul-NITD/Poker"
)

func main() {
	storage := poker.NewInMemoryStorage()
	server := poker.NewServer(&storage)
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(server.ServeHTTP)))
}
