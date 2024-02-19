package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Panic("Error opening database connection: ", err)
	}

	return db, nil
}
