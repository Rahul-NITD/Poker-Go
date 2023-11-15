package poker

type InMemoryStorage struct {
	store map[string]int
}

func (str *InMemoryStorage) GetScore(player string) (int, error) {
	score, ok := str.store[player]
	if !ok {
		return 0, ERRORPlayerNotFound
	}
	return score, nil
}

func NewInMemoryStorage() InMemoryStorage {
	return InMemoryStorage{
		store: map[string]int{
			"dev": 5,
		},
	}
}
