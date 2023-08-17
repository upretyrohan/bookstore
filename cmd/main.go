package main

import (
	"log"
	"net/http"
	"github.com/upretyrohan/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/upretyrohan/bookstore/pkg/database"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	database.Init()
	log.Fatal(http.ListenAndServe("localhost:9010",r))
}