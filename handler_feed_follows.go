package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aaradhyakul/rssagg/internal/databases"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)
func (apiCfg *apiConfig)handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user databases.User){
	type parameters struct{
		FeedId uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create feed follow",err))
		return
	}
	feedFollow,err := apiCfg.DB.CreateFeedFollow(r.Context(), databases.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),		
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: params.FeedId,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create feed:%v",err))
		return
	}
	respondWithJSON(w,201,databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig)handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user databases.User){
	feedFollows,err := apiCfg.DB.GetFeedFollows(r.Context(),user.ID)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't get feed follows:%v",err))
		return
	}
	respondWithJSON(w,201,databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig)handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user databases.User){
	feedFollowIDStr := chi.URLParam(r,"feedFollowID")
	feedFollowID,err := uuid.Parse(feedFollowIDStr)
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't parse feed follow ID: %v",err))
		return;
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(),databases.DeleteFeedFollowParams{
		ID:feedFollowID,
		UserID: user.ID,
	})
	if err!= nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't delete feed follow: %v",err))
		return
	}
	respondWithJSON(w,200,struct{}{})
}
