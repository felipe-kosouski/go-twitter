package main

import (
	"context"
	"go_twitter/internal/handler"
	"go_twitter/internal/repository"
	"go_twitter/internal/service"
	"go_twitter/pkg/db"
	"log"
	"net/http"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	userRepo := repository.NewUserRepository(conn)
	tweetRepo := repository.NewTweetRepository(conn)

	userService := service.NewUserService(userRepo)
	tweetService := service.NewTweetService(tweetRepo)

	userHandler := handler.NewUserHandler(userService)
	tweetHandler := handler.NewTweetHandler(tweetService)

	r := handler.NewRouter(userHandler, tweetHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
