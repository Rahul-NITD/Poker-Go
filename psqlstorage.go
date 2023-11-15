package poker

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "gotestdb"
)

func ConnectDB() (*sql.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(db *sql.DB, temp bool) error {

	var temptable string
	if temp {
		temptable = "temp"
	}

	query := fmt.Sprintf(`Create %s table if not exists League (
    name varchar(25) PRIMARY KEY,
    wins integer DEFAULT 0
  );`, temptable)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil

}

func InsertPlayer(db *sql.DB, player Player) error {
	query := `INSERT INTO League (name, wins) values ($1, $2)`
	_, err := db.Exec(query, player.Name, player.Wins)

	if err != nil {
		return err
	}
	return nil
}

func ReadLeague(db *sql.DB) ([]Player, error) {

	query := `SELECT name, wins from League ORDER BY wins desc`
	rows, err := db.Query(query)
	if err != nil {
		return []Player{}, nil
	}
	defer rows.Close()

	var name string
	var wins int
	res := []Player{}
	for rows.Next() {
		err := rows.Scan(&name, &wins)
		if err != nil {
			return []Player{}, nil
		}
		res = append(res, Player{Name: name, Wins: wins})
	}
	return res, nil
}

func UpdateValues(db *sql.DB, playerName string) error {

	score, _ := GetScore(db, playerName)
	if score == 0 {
		InsertPlayer(db, Player{
			Name: playerName,
			Wins: 0,
		})
	}

	query := "UPDATE League SET wins=wins+1 where name=$1"
	_, err := db.Exec(query, playerName)
	return err
}

func GetScore(db *sql.DB, playerName string) (int, error) {
	query := "SELECT wins from League where name=$1"
	row := db.QueryRow(query, playerName)
	var score int
	row.Scan(&score)
	if score == 0 {
		return 0, errors.New("Player Does not Exist")
	}
	return score, nil
}

type DBStorage struct {
	db *sql.DB
}

func (storage *DBStorage) GetScore(player string) (int, error) {

	return GetScore(storage.db, player)

}

func (storage *DBStorage) RecordWin(player string) error {
	return UpdateValues(storage.db, player)
}

func (storage *DBStorage) GetLeague() []Player {
	res, _ := ReadLeague(storage.db)
	return res
}

func NewDBStorage(test bool) (*DBStorage, func(str *DBStorage), error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, func(str *DBStorage) {}, err
	}
	if err := CreateTable(db, test); err != nil {
		return nil, func(str *DBStorage) {}, err
	}

	disconnect := func(str *DBStorage) {
		log.Println("Closing DB")
		str.db.Close()
	}

	return &DBStorage{
		db: db,
	}, disconnect, nil

}
