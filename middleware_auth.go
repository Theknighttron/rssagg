package main

import (
    "net/http"
    "fmt"
    "github.com/TheKnighttron/rssagg/internal/database"
    auth "github.com/TheKnighttron/rssagg/internal"
)


type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        apiKey, err := auth.GetAPIKey(r.Header)
        if err != nil {
            respondWithError(w, 403, fmt.Sprintf("Authentication error %v", err))
            return
        }


        user, err :=  apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
        if err != nil {
            respondWithError(w, 400, fmt.Sprintf("Couldn't get user %v", err))
            return
        }

        handler(w, r, user)
    }
}
