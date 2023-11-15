package poker_test

import (
	"net/http"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestHTTPServer(t *testing.T) {

	t.Run("Test GET route", func(t *testing.T) {

		storage := NewSTUBStorage()
		server := poker.NewServer(&storage)

		cases := []struct {
			title            string
			path             string
			expectedCode     int
			expectedResponse string
		}{
			{
				"Server Listens dev",
				"/players/dev",
				http.StatusOK,
				"1",
			},
			{
				"Server listens Rahul",
				"/players/Rahul",
				http.StatusOK,
				"2",
			},
			{
				"Server listens Akku",
				"/players/Akku",
				http.StatusOK,
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

	t.Run("Test POST route", func(t *testing.T) {
		storage := NewSTUBStorage()
		server := poker.NewServer(&storage)

		cases := []struct {
			title            string
			path             string
			expectedCode     int
			expectedResponse string
		}{
			{
				"Server Records dev",
				"/players/dev",
				http.StatusAccepted,
				"2",
			},
			{
				"Server Records Akku",
				"/players/Akku",
				http.StatusAccepted,
				"4",
			},
			{
				"Server Records a New Player",
				"/players/IAmUndefined",
				http.StatusAccepted,
				"1",
			},
		}

		for _, test := range cases {
			t.Run(test.title, func(t *testing.T) {
				res, req := CreatePostRequest(test.path)
				server.ServeHTTP(res, req)
				AssertStatusCode(t, res.Code, test.expectedCode)
				res, req = CreateGetRequest(test.path)
				server.ServeHTTP(res, req)
				AssertResponseBody(t, res.Body.String(), test.expectedResponse)
			})
		}

	})
}
