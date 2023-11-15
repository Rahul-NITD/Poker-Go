package poker_test

import (
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestInMemoryStorage(t *testing.T) {
	storage := poker.NewInMemoryStorage()

	t.Run("Test Existing Player Dev", func(t *testing.T) {
		got, err := storage.GetScore("dev")
		AssertNilError(t, err)
		AssertScores(t, got, 5)
	})

	t.Run("Test Not Existing Player Rahul", func(t *testing.T) {
		got, err := storage.GetScore("Rahul")
		AssertError(t, err, poker.ERRORPlayerNotFound)
		AssertScores(t, got, 0)
	})

	t.Run("Test Recording existing dev", func(t *testing.T) {
		want, _ := storage.GetScore("dev")
		want++
		err := storage.RecordWin("dev")
		AssertNilError(t, err)
		got, _ := storage.GetScore("dev")
		AssertScores(t, got, want)
	})

	t.Run("Test Recording new player", func(t *testing.T) {
		player := "IAmUndefined"
		want, _ := storage.GetScore(player)
		want++
		err := storage.RecordWin(player)
		AssertNilError(t, err)
		got, _ := storage.GetScore(player)
		AssertScores(t, got, want)
	})

	storage = poker.NewInMemoryStorage()
	storage.RecordWin("TempBoi")

	t.Run("Test League", func(t *testing.T) {
		got := storage.GetLeague()
		want := []poker.Player{
			{
				Name: "dev",
				Wins: 5,
			},
			{
				Name: "TempBoi",
				Wins: 1,
			},
		}
		AssertLeague(t, got, want)
	})

}
