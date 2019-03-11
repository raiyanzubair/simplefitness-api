package main

import (
	_ "github.com/lib/pq"
	_ "github.com/rubenv/sql-migrate"
	"log"
	"net/http"
	"os"
	"simplefitnessApi/delivery/http/app"
)

func main() {
	log.Print("Running")

	myApp, _ := app.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, myApp.Router)
}
