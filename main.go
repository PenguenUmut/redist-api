package main

import (
	//	"encoding/json"
	"log"
	"net/http"

	"./handler"
	// "github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hello Redist-Api!")

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/users", handler.Home)
	r.HandleFunc("/articles", handler.Home)
	http.Handle("/", r)
	http.ListenAndServe(":9000", r)
}
