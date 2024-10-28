package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go_twitter/internal/models"
)

type TweetRepository interface {
	GetTweets(ctx context.Context) ([]*models.Tweet, error)
	GetTweetById(ctx context.Context, id int) (*models.Tweet, error)
	CreateTweet(ctx context.Context, tweet *models.Tweet) error
	UpdateTweet(ctx context.Context, tweet *models.Tweet) error
	DeleteTweet(ctx context.Context, id int) error
}

type tweetRepository struct {
	db *pgx.Conn
}

func NewTweetRepository(db *pgx.Conn) TweetRepository {
	return &tweetRepository{db: db}
}

func (t tweetRepository) GetTweets(ctx context.Context) ([]*models.Tweet, error) {
	rows, err := t.db.Query(ctx, "SELECT id, user_id, content, created_at FROM tweets")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tweets []*models.Tweet
	for rows.Next() {
		tweet := &models.Tweet{}
		err := rows.Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, tweet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tweets, nil
}

func (t tweetRepository) GetTweetById(ctx context.Context, id int) (*models.Tweet, error) {
	tweet := &models.Tweet{}
	err := t.db.QueryRow(ctx, "SELECT id, user_id, content, created_at FROM tweets WHERE id=$1", id).Scan(&tweet.ID, &tweet.UserID, &tweet.Content, &tweet.CreatedAt)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t tweetRepository) CreateTweet(ctx context.Context, tweet *models.Tweet) error {
	_, err := t.db.Exec(ctx, "INSERT INTO tweets (user_id, content) VALUES ($1, $2)", tweet.UserID, tweet.Content)
	return err
}

func (t tweetRepository) UpdateTweet(ctx context.Context, tweet *models.Tweet) error {
	_, err := t.db.Exec(ctx, "UPDATE tweets SET content=$1 WHERE id=$2", tweet.Content, tweet.ID)
	return err
}

func (t tweetRepository) DeleteTweet(ctx context.Context, id int) error {
	_, err := t.db.Exec(ctx, "DELETE FROM tweets WHERE id=$1", id)
	return err
}
