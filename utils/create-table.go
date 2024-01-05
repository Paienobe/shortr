package utils

import (
	"database/sql"
	"log"

	"github.com/Paienobe/go-url-shortener/queries"
)

func CreateTable(db *sql.DB) {
	query := queries.CreateTableQuery
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
