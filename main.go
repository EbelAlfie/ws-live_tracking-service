package main

import (
	"fmt"
	"live-track/main/route"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	route.LocationRoute(mux)

	fmt.Print("Hello world")

	err := http.ListenAndServe("localhost:3069", mux)
	if err != nil {
		log.Fatal("ERROR", err)
	}
}
