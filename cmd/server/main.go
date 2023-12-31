package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"web-server/pkg/api"
	"web-server/pkg/db"
)

func main() {
	log.Print("server has started")

	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error starting the database %v", err)
	}

	router := api.StartAPI(pgdb)

	port := os.Getenv("PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
	}
}
