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
}
