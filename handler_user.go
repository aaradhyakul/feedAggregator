package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aaradhyakul/rssagg/internal/databases"
	"github.com/google/uuid"
)
func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error Parsing JSON:%v",err))
		return
	}
	user,err := apiCfg.DB.CreateUser(r.Context(), databases.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user:%v",err))
		return
	}
	respondWithJSON(w,201,databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user databases.User){
	respondWithJSON(w,200,databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user databases.User){
	posts,err := apiCfg.DB.GetPostsForUser(r.Context(),databases.GetPostsForUserParams{
		UserID:user.ID,
		Limit: 10,
	})
	if err!=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't get posts: %v",err))
	}
	respondWithJSON(w,200,databasePostsToPosts(posts))
}
