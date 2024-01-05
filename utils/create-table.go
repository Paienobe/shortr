package utils

import (
	"database/sql"
	"log"

	"github.com/Paienobe/go-url-shortener/constants"
)

func CreateTable(db *sql.DB) {
	query := constants.CreateTableQuery
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
