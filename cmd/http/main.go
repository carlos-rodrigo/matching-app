package main

import (
	"log"
	"net/http"

	delivery "github.com/carlos-rodrigo/matching-app/pkg/delivery/http"
)

func main() {
	router := delivery.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
