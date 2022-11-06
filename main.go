package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", Routes())
	if err != nil {
		log.Fatal(err)
	}
}
