package handler

import (
	// "encoding/json"
	"log"
	"net/http"

	"../db"
	"github.com/gorilla/mux"
)

func Home(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Println("Request:", req.Method, req.RequestURI, vars)
	Send(&rw, db.Get(req.Method, req.RequestURI, vars))
}
