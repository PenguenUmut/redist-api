package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Send(rw *http.ResponseWriter, v interface{}) {
	(*rw).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*rw).Header().Set("Access-Control-Allow-Origin", "*")
	(*rw).Header().Set("Cache-Control", "no-cache")
	(*rw).WriteHeader(http.StatusOK)
	json.NewEncoder(*rw).Encode(v)
	log.Println("Response", v)
}
