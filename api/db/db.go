package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../data/database.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetChurnUserScore(idUser string, con *sql.DB) (float64, error) {

	score := -1.

	query := `
	SELECT nrProbaRank
	FROM churn_score
	WHERE idCliente = ?`

	state, err := con.Prepare(query)
	if err != nil {
		return score, err
	}

	rows, err := state.Query(idUser)
	if err != nil {
		return score, err
	}

	for rows.Next() {
		rows.Scan(&score)
	}

	return score, nil
}
