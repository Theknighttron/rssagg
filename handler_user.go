package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/TheKnighttron/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type paramaters struct {
        Name string `json:"name"`
    }

    params := paramaters{}

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error parsinf JSON %v", err))
        return
    }

    user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
        ID: uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        Name: params.Name,
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't create user %v", err))
        return
    }

    respondWithJSON(w, 200, databaseUserToUser(user))
}
