package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Paienobe/go-url-shortener/queries"
	"github.com/Paienobe/go-url-shortener/utils"
	"github.com/gorilla/mux"
)

func VisitShort(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.EnableCors(&w)
	vars := mux.Vars(r)
	shortKeyParam := vars["short_key"]
	var originalUrl string
	query := queries.FindLinkQuery
	err := db.QueryRow(query, shortKeyParam).Scan(&originalUrl)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, originalUrl, http.StatusFound)
}
