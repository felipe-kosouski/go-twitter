package service

import (
	"context"
	"go_twitter/internal/models"
	"go_twitter/internal/repository"
)

type tweetService struct {
	tweetRepo repository.TweetRepository
}

func NewTweetService(tweetRepo repository.TweetRepository) TweetService {
	return &tweetService{tweetRepo: tweetRepo}
}

func (t tweetService) GetTweets(ctx context.Context) ([]*models.Tweet, error) {
	return t.tweetRepo.GetTweets(ctx)
}

func (t tweetService) GetTweetById(ctx context.Context, id int) (*models.Tweet, error) {
	return t.tweetRepo.GetTweetById(ctx, id)
}

func (t tweetService) CreateTweet(ctx context.Context, tweet *models.Tweet) error {
	return t.tweetRepo.CreateTweet(ctx, tweet)
}

func (t tweetService) UpdateTweet(ctx context.Context, tweet *models.Tweet) error {
	return t.tweetRepo.UpdateTweet(ctx, tweet)
}

func (t tweetService) DeleteTweet(ctx context.Context, id int) error {
	return t.tweetRepo.DeleteTweet(ctx, id)
}
