package main

import (
	"log"
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 { // 500: bug not on client end
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error`
	}

	respondWithJSON(w, code, errResponse{msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to marshal payload: %v", payload)
		// internal service error
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) //use passed in code
	w.Write(dat)
}