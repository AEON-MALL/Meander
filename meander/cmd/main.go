package main

import (
	"encoding/json"
	"fmt"
	"meander/meander"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	http.HandleFunc("/journeys",func(w http.ResponseWriter, r *http.Request){
		respond(w, r, meander.Journeys)
	})
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func loadEnv(){
	err := godotenv.Load(".env")
	if err != nil{
		fmt.Printf("読み込みできませんでした: %v", err)
	}
	meander.APIKey = os.Getenv("googleplaceapikey")
}

func respond(w http.ResponseWriter, r *http.Request, data []any) error {
	publicData := make([]any, len(data))
	for i, d := range data{
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}



