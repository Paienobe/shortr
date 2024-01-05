package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Paienobe/go-url-shortener/queries"
	"github.com/gorilla/mux"
)

func DeleteLink(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	shortKeyParam := vars["short_key"]
	query := queries.DeleteLinkQuery

	_, err := db.Exec(query, shortKeyParam)
	if err != nil {
		log.Fatal(err)
	}
}
