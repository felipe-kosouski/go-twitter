package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(userHandler *UserHandler, tweetHandler *TweetHandler) *mux.Router {
	r := mux.NewRouter()
	r.Use(jsonMiddleware)

	r.HandleFunc("/users", userHandler.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserById).Methods(http.MethodGet)
	r.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods(http.MethodDelete)

	r.HandleFunc("/tweets", tweetHandler.GetTweets).Methods(http.MethodGet)
	r.HandleFunc("/tweets/{id:[0-9]+}", tweetHandler.GetTweetById).Methods(http.MethodGet)
	r.HandleFunc("/tweets", tweetHandler.CreateTweet).Methods(http.MethodPost)
	r.HandleFunc("/tweets/{id:[0-9]+}", tweetHandler.UpdateTweet).Methods(http.MethodPut)
	r.HandleFunc("/tweets/{id:[0-9]+}", tweetHandler.DeleteTweet).Methods(http.MethodDelete)

	return r
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
