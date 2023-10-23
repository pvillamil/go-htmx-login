package database

import (
	"database/sql"
	"log"
)

func ConnectToSqlLite(databasePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

type User struct {
	Id        int
	Username  string
	Password  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
