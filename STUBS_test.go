package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

// ASSERTIONS
func AssertStatusCode(t testing.TB, resCode, wantedCode int) {
	t.Helper()
	if resCode != wantedCode {
		t.Errorf("Wanted code %d, got code %d", wantedCode, resCode)
	}
}

func AssertResponseBody(t testing.TB, response, wanted string) {
	t.Helper()
	if response != wanted {
		t.Errorf("Got Score %q, wanted score %q", response, wanted)
	}
}

func AssertNilError(t testing.TB, got error) {
	t.Helper()
	AssertError(t, got, nil)
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q, got %q", got.Error(), want.Error())
	}
}

func AssertScores(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected score %d, got %d", want, got)
	}
}

// REQUESTS
func CreateGetRequest(path string) (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	res := httptest.NewRecorder()
	return res, req
}

// STUBStorage
type STUBStorage struct {
	Scores map[string]int
}

func (str *STUBStorage) GetScore(player string) (int, error) {
	score, ok := str.Scores[player]
	if !ok {
		return 0, poker.ERRORPlayerNotFound
	}
	return score, nil
}

func NewSTUBStorage() STUBStorage {
	return STUBStorage{
		Scores: map[string]int{
			"Rahul": 2,
			"Akku":  3,
			"dev":   1,
		},
	}
}
