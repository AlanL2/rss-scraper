package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { //always takes ResponseWriter as first param, http.Request as second
	respondWithJSON(w, 200, struct{}{})
}