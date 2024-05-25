package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
