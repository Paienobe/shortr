package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/go-url-shortener/constants"
	"github.com/Paienobe/go-url-shortener/dto"
	"github.com/Paienobe/go-url-shortener/types"
	"github.com/Paienobe/go-url-shortener/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("dbUrl does not exist in environment")
	}

	portString := ":" + os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Portstring does not exist in environment")
	}

	dbConn, err := sql.Open("postgres", dbUrl)
	defer dbConn.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}

	CreateTable(dbConn)

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { CreateLink(w, r, dbConn) }).Methods("POST")

	fmt.Printf("Server listening on PORT: %s\n", portString)
	log.Fatal(http.ListenAndServe(portString, router))

}

func CreateTable(db *sql.DB) {
	query := constants.CreateTableQuery
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertLinkRow(db *sql.DB, link types.Link) types.Link {
	query := constants.CreateLinkQuery
	var longUrl string
	var shortKey string
	var shortUrl string
	err := db.QueryRow(query, link.LongUrl, link.ShortKey, link.ShortUrl).Scan(&longUrl, &shortKey, &shortUrl)
	if err != nil {
		fmt.Println("err here")
		log.Fatal(err)
	}
	return types.Link{LongUrl: longUrl, ShortKey: shortKey, ShortUrl: shortUrl}
}

func CreateLink(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// use r.TLS to toggle http and https
	fmt.Println(r.TLS, "tls")

	var requestBody dto.CreateLinkDto

	if err := json.Unmarshal(rawBody, &requestBody); err != nil {
		log.Fatal("Error parsing data")
	}

	fmt.Println(r.Host)

	longUrl := requestBody.Url
	shortKey := utils.GenerateHash(requestBody.Url)
	shortLink := fmt.Sprintf("%s/%s", r.Host, shortKey)

	generatedLink := types.Link{LongUrl: longUrl, ShortKey: shortKey, ShortUrl: shortLink}

	shortr := InsertLinkRow(db, generatedLink)
	res, err := json.Marshal(shortr)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
