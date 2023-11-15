package poker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PokerStorage interface {
	GetScore(player string) (int, error)
	RecordWin(player string) error
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

type PokerServer struct {
	ScoreStorage PokerStorage
	http.Handler
}

func NewServer(storage PokerStorage) *PokerServer {
	server := new(PokerServer)
	server.ScoreStorage = storage
	router := http.NewServeMux()
	router.Handle("/players/", http.HandlerFunc(server.playersRouteHandler))
	router.Handle("/league", http.HandlerFunc(server.leagueRouteHandler))
	server.Handler = router
	return server
}

func (server *PokerServer) playersRouteHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		getScore(w, player, server.ScoreStorage)
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		recordWin(w, player, server.ScoreStorage)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (server *PokerServer) leagueRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getLeague(w, server.ScoreStorage)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getLeague(w http.ResponseWriter, storage PokerStorage) {
	playersLeague := storage.GetLeague()
	b := bytes.Buffer{}
	json.NewEncoder(&b).Encode(playersLeague)
	w.Write(b.Bytes())
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

func recordWin(w http.ResponseWriter, player string, storage PokerStorage) {
	err := storage.RecordWin(player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
