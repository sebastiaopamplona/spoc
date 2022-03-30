package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"net/http"
	"simaopoc/api"
	"time"
)

func main() {
	// init route https://github.com/gorilla/muxr
	router := mux.NewRouter()

	// init db connection
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		"localhost", "5433", "postgres", "simao", "secret")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to DB, %s", err))
	}

	service := api.New(router, db)

	// define endpoints ("endpoint", "function that handles it")
	router.HandleFunc("/fetchdata", service.FetchData)
	router.HandleFunc("/createrandomuser", service.CreateRandomUser)

	// create server
	srv := &http.Server{
		Handler: router,
		Addr:    ":8090",
		// good pracice
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start server
	fmt.Println("listening on localhost:8090...")
	log.Fatal(srv.ListenAndServe())
}
