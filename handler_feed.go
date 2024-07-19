package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aaradhyakul/rssagg/internal/databases"
	"github.com/google/uuid"
)
func (apiCfg *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request, user databases.User){
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error Parsing JSON:%v",err))
		return
	}
	feed,err := apiCfg.DB.CreateFeed(r.Context(), databases.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create feed:%v",err))
		return
	}
	respondWithJSON(w,201,databaseFeedToFeed(feed))
}
func (apiCfg *apiConfig)handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds,err := apiCfg.DB.GetFeeds(r.Context())
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't fetch feeds:%v",err))
		return
	}
	respondWithJSON(w,201,databaseFeedsToFeeds(feeds))
}