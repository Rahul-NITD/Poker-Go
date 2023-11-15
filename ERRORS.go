package poker

type PokerError struct {
	string
}

func (err PokerError) Error() string {
	return err.string
}

var ERRORPlayerNotFound = PokerError{"could not find player"}
