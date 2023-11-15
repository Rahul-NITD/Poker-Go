package poker

import "sort"

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

func (str *InMemoryStorage) RecordWin(player string) error {
	str.store[player]++
	return nil
}

func (str *InMemoryStorage) GetLeague() []Player {
	var res []Player
	for key, value := range str.store {
		res = append(res, Player{
			Name: key,
			Wins: value,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Wins > res[j].Wins
	})

	return res
}

func NewInMemoryStorage() InMemoryStorage {
	return InMemoryStorage{
		store: map[string]int{
			"dev": 5,
		},
	}
}
