package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/alexaaaant/user-service/internal/http"
	"github.com/alexaaaant/user-service/internal/repository"
	"github.com/alexaaaant/user-service/internal/service"
)

func main() {
	// --- wiring ---
	userRepo := repository.NewUserMemoryRepo()
	userService := service.NewUserService(userRepo)
	handler := httpHandler.NewHandler(userService)
	router := httpHandler.NewRouter(handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// --- запуск сервера ---
	go func() {
		log.Println("server started on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	// --- ожидание сигнала ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("shutting down server...")

	// --- graceful shutdown ---
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	} else {
		log.Println("server stopped gracefully")
	}
}
