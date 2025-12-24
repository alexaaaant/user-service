package main

import (
	"log"
	"net/http"

	httpHandler "github.com/alexaaaant/user-service/internal/http"
	"github.com/alexaaaant/user-service/internal/repository"
	"github.com/alexaaaant/user-service/internal/service"
)

func main() {
	// repository
	userRepo := repository.NewUserMemoryRepo()

	// service
	userService := service.NewUserService(userRepo)

	handler := httpHandler.NewHandler(userService)
	router := httpHandler.NewRouter(handler)

	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
