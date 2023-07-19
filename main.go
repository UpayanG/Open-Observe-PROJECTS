package main

import (
	"encoding/json"
	"net/http"
)

func getWorld(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Word string `json:"hello"`
	}{
		Word: "world",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", getWorld)
	http.ListenAndServe(":8081", nil)
}
