package db

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatal(err)
	}
}