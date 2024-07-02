package db

import (
	"fmt"
	"database/sql"
)

func PostgresConnection() *sql.DB {

	connection := "Get your own connection string"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		fmt.Printf("Error opening database. %s", err.Error())
	}

	db.Ping()

	if err != nil {panic(err.Error())}

	return db
}