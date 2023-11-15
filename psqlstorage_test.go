package poker_test

import (
	"reflect"
	"testing"

	poker "github.com/Rahul-NITD/Poker"
)

func TestPSQLStorage(t *testing.T) {

	db, err := poker.ConnectDB()
	if err != nil {
		t.Fatalf("Could not connect to DB, %q", err.Error())
	}
	defer db.Close()

	t.Run("test creating table", func(t *testing.T) {

		err := poker.CreateTable(db, true)

		if err != nil {
			t.Fatalf("Error creating table, %q", err)
		}

	})

	t.Run("Test inserting into table", func(t *testing.T) {

		player := poker.Player{
			Name: "Rahul",
			Wins: 4,
		}

		err := poker.InsertPlayer(db, player)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Read League", func(t *testing.T) {

		poker.InsertPlayer(db, poker.Player{
			Name: "Akku",
			Wins: 5,
		})
		want := []poker.Player{
			{
				Name: "Akku",
				Wins: 5,
			},
			{
				Name: "Rahul",
				Wins: 4,
			},
		}

		got, err := poker.ReadLeague(db)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Got %v != want %v", got, want)
		}

		if err != nil {
			t.Fatalf("Error occured, %q", err.Error())
		}

	})

	t.Run("Test Update values", func(t *testing.T) {
		poker.UpdateValues(db, "Akku")
		want := []poker.Player{
			{
				Name: "Akku",
				Wins: 6,
			},
			{
				Name: "Rahul",
				Wins: 4,
			},
		}

		got, err := poker.ReadLeague(db)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

		if err != nil {
			t.Error(err.Error())
		}

	})

	t.Run("Test get score", func(t *testing.T) {
		score, err := poker.GetScore(db, "Rahul")
		if err != nil {
			t.Fatal(err)
		}
		if score != 4 {
			t.Errorf("Got %d, want %d", score, 4)
		}
	})

}

func TestIntegration(t *testing.T) {

	storage, close, err := poker.NewDBStorage(true)
	if err != nil {
		t.Errorf("Error occured, %q", err.Error())
	}
	defer close(storage)

	t.Run("Test Interface", func(t *testing.T) {

		_, err := storage.GetScore("Rahul")

		if err.Error() != "Player Does not Exist" {
			t.Fatalf("This is a new error, err : %q", err.Error())
		}

	})

	t.Run("Test Record Win", func(t *testing.T) {
		err := storage.RecordWin("Rahul")
		if err != nil {
			t.Fatal(err)
		}
		score, err := storage.GetScore("Rahul")
		if err != nil {
			t.Fatal(err)
		}
		if score != 1 {
			t.Errorf("got %d, want %d", score, 1)
		}
	})

	for i := 0; i < 6; i++ {
		storage.RecordWin("Akku")
	}

	t.Run("Test League", func(t *testing.T) {
		got := storage.GetLeague()
		want := []poker.Player{
			{
				Name: "Akku",
				Wins: 6,
			},
			{
				Name: "Rahul",
				Wins: 1,
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
