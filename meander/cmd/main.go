package main

import (
	"encoding/json"
	_ "log"
	"meander/meander"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv"
)

func main() {
	/*&if err := loadEnv(); err != nil {
		return
	}*/

	meander.APIKey = "AIzaSyAZEZwUbm4QBz_ODRmgA4oHtbVdHZca4GE"

	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	}))
	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		var err error
		if q.Lat, err = strconv.ParseFloat(r.URL.Query().Get("lat"), 64); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if q.Lng, err = strconv.ParseFloat(r.URL.Query().Get("lng"), 64); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if q.Radius, err = strconv.Atoi(r.URL.Query().Get("radius")); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		respond(w, r, places)
	}))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

/*
func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("unable read: %v", err)
	}
	meander.APIKey = os.Getenv("googleplaceapikey")
	return nil
}
*/

func respond(w http.ResponseWriter, r *http.Request, data []any) error {
	publicData := make([]any, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}
