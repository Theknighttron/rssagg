package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/TheKnighttron/rssagg/internal/database"
    "github.com/google/uuid"
	"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    type paramaters struct {
        FeedID uuid.UUID `json:"feed_id"`
    }

    params := paramaters{}

    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Error parsinf JSON %v", err))
        return
    }

    feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
        ID: uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        UserID: user.ID,
        FeedID: params.FeedID,
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow %v", err))
        return
    }

    respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow %v", err))
        return
    }

    respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))
}


func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollowString := chi.URLParam(r, "feedFollowID")
    feedFollowID, err := uuid.Parse(feedFollowString)
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id %v", err))
        return
    }

    err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
        ID: feedFollowID,
        UserID: user.ID,
    })
    if err != nil {
        respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow %v", err))
        return
    }

    respondWithJSON(w, 200, struct{}{})
}

