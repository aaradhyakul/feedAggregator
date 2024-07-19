package main

import (
	"fmt"
	"net/http"

	"github.com/aaradhyakul/rssagg/internal/auth"
	"github.com/aaradhyakul/rssagg/internal/databases"
)
type authedHandler 	func(http.ResponseWriter, *http.Request, databases.User)
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKey(r.Header)
		if err!= nil{
			respondWithError(w,403,fmt.Sprintf("Auth Error: %v",err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err!=nil{
			respondWithError(w,400,fmt.Sprintf("Coundn't get user: %v",err))
			return
		}
		handler(w,r,user)
	}
}