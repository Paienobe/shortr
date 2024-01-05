package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Paienobe/go-url-shortener/dto"
	"github.com/Paienobe/go-url-shortener/types"
	"github.com/Paienobe/go-url-shortener/utils"
)

func CreateLink(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	utils.EnableCors(&w)
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var transferProtocol string

	if r.TLS == nil {
		transferProtocol = "http://"
	} else {
		transferProtocol = "https://"
	}

	var requestBody dto.CreateLinkDto

	if err := json.Unmarshal(rawBody, &requestBody); err != nil {
		http.Error(w, "Error unmarshalling data", http.StatusBadRequest)
	}

	longUrl := requestBody.Url
	shortKey := utils.GenerateHash(requestBody.Url)
	shortLink := fmt.Sprintf("%s%s/%s", transferProtocol, r.Host, shortKey)

	generatedLink := types.Link{LongUrl: longUrl, ShortKey: shortKey, ShortUrl: shortLink}

	shortr := utils.InsertLinkRow(db, generatedLink)
	res, err := json.Marshal(shortr)
	if err != nil {
		http.Error(w, "Error marshalling data", http.StatusBadRequest)
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
