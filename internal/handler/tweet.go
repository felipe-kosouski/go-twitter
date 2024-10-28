package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go_twitter/internal/models"
	"go_twitter/internal/service"
	"net/http"
	"strconv"
)

type TweetHandler struct {
	tweetService service.TweetService
}

func NewTweetHandler(tweetService service.TweetService) *TweetHandler {
	return &TweetHandler{tweetService: tweetService}
}

func (t *TweetHandler) GetTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := t.tweetService.GetTweets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tweets)
}

func (t *TweetHandler) GetTweetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	tweet, err := t.tweetService.GetTweetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tweet)
}

func (t *TweetHandler) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := t.tweetService.CreateTweet(r.Context(), &tweet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (t *TweetHandler) UpdateTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	if err := json.NewDecoder(r.Body).Decode(&tweet); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := t.tweetService.UpdateTweet(r.Context(), &tweet); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (t *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid tweet ID", http.StatusBadRequest)
		return
	}
	if err := t.tweetService.DeleteTweet(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
