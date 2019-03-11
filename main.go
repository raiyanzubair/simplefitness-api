package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"simplefitnessApi/delivery/http/app"
)

func main() {
	log.Print("Running")

	myApp, _ := app.New()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, myApp.Router)
}
