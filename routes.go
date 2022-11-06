package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() http.Handler {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", HomeHandler)
	return r
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
