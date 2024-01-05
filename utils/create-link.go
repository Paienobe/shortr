package utils

import (
	"database/sql"
	"log"

	"github.com/Paienobe/go-url-shortener/queries"
	"github.com/Paienobe/go-url-shortener/types"
)

func InsertLinkRow(db *sql.DB, link types.Link) types.Link {
	query := queries.CreateLinkQuery
	var longUrl string
	var shortKey string
	var shortUrl string
	err := db.QueryRow(query, link.LongUrl, link.ShortKey, link.ShortUrl).Scan(&longUrl, &shortKey, &shortUrl)
	if err != nil {
		log.Fatal(err)
	}
	return types.Link{LongUrl: longUrl, ShortKey: shortKey, ShortUrl: shortUrl}
}
