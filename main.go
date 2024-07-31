package main

import (
	"fmt"
	"live-track/main/route"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	baseAddr := os.Getenv("BASE_ADDR")

	mux := http.NewServeMux()

	route.LocationRoute(mux)

	fmt.Print("Hello world")

	err = http.ListenAndServe(baseAddr, mux)
	if err != nil {
		log.Fatal("ERROR", err)
	}
}
