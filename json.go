package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string){
    if statusCode > 499 {
        log.Println("Responding with 500 error: ", msg)
    }

    type errorResponse struct {
        Error string `json:"error"`
    }

    respondWithJSON(w, statusCode, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
    // Marshall the given data to json string
    data, err := json.Marshal(payload)
    if err != nil {
        log.Printf("Failed to marshall JSON response %v", payload)
        w.WriteHeader(500)
        return
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(data)
}
