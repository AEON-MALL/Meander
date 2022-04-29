package main

import (
	"encoding/json"
	"meander/meander"
	"net/http"
)

func main() {
	http.HandleFunc("/journeys",func(w http.ResponseWriter, r *http.Request){
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080",http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []any) error {
	return json.NewEncoder(w).Encode(data)
}