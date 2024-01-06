package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/go-url-shortener/controllers"
	"github.com/Paienobe/go-url-shortener/utils"
	"github.com/go-chi/cors"
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

	utils.CreateTable(dbConn)

	router := mux.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://shortr-ui.vercel.app/", "http://localhost:5173/"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.HandleFunc("/create-link", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateLink(w, r, dbConn)
	}).Methods("POST")

	router.HandleFunc("/{short_key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.VisitShort(w, r, dbConn)
	}).Methods("GET")

	router.HandleFunc("/{short_key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteLink(w, r, dbConn)
	}).Methods("DELETE")

	fmt.Printf("Server listening on PORT: %s\n", portString)
	log.Fatal(http.ListenAndServe(portString, router))

}
