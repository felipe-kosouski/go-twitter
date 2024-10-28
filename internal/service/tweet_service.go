package service

import (
	"context"
	"go_twitter/internal/models"
)

type TweetService interface {
	GetTweets(ctx context.Context) ([]*models.Tweet, error)
	GetTweetById(ctx context.Context, id int) (*models.Tweet, error)
	CreateTweet(ctx context.Context, tweet *models.Tweet) error
	UpdateTweet(ctx context.Context, tweet *models.Tweet) error
	DeleteTweet(ctx context.Context, id int) error
}
