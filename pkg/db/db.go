package db

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/go_twitter")
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}
	return conn, nil
}
