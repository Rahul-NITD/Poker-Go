package poker

import (
	"fmt"
	"net/http"
	"strings"
)

type PokerStorage interface {
	GetScore(player string) (int, error)
}

type PokerServer struct {
	ScoreStorage PokerStorage
}

func (server *PokerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")
	getScore(w, player, server.ScoreStorage)

}

func getScore(w http.ResponseWriter, player string, storage PokerStorage) {
	score, err := storage.GetScore(player)
	switch err {
	case nil:
		fmt.Fprintf(w, "%d", score)
	case ERRORPlayerNotFound:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Player Not Found")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Player Not Found")
	}
}
