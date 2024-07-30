package main

import (
	"fmt"
	"live-track/main/route"
	"log"
	"net/http"
)

func main() {

	route := route.LocationRoute()

	fmt.Print("Hello world")

	err := http.ListenAndServe("localhost:3069", route)
	if err != nil {
		log.Fatal("ERROR", err)
	}
}
