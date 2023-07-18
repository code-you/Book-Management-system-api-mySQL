package main

import (
	"github.com/code-you/Book-Management-system-api-mySQL/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting")
	r := mux.NewRouter()
	routes.RegisterBookStoreRoute(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
