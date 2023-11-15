package poker_test

import (
	"net/http"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestHTTPServer(t *testing.T) {

	t.Run("Test Server listens", func(t *testing.T) {

		storage := NewSTUBStorage()
		server := poker.PokerServer{&storage}

		cases := []struct {
			title            string
			path             string
			expectedCode     int
			expectedResponse string
		}{
			{
				"Server Listens dev",
				"/players/dev",
				200,
				"1",
			},
			{
				"Server listens Rahul",
				"/players/Rahul",
				200,
				"2",
			},
			{
				"Server listens Akku",
				"/players/Akku",
				200,
				"3",
			},
			{
				"Server listens undefined player",
				"/players/IAmUndefined",
				http.StatusNotFound,
				"Player Not Found",
			},
		}

		for _, test := range cases {
			t.Run(test.title, func(t *testing.T) {
				res, req := CreateGetRequest(test.path)
				server.ServeHTTP(res, req)
				AssertStatusCode(t, res.Code, test.expectedCode)
				AssertResponseBody(t, res.Body.String(), test.expectedResponse)
			})
		}

	})
}
