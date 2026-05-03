package main

import (
	"context"
	"log"

	"github.com/umarov-safar/user-management-api/internal/config"
	"github.com/umarov-safar/user-management-api/internal/database"
	"github.com/umarov-safar/user-management-api/internal/repositories"
	"github.com/umarov-safar/user-management-api/internal/services"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("Starting server on %s:%s", cfg.ServerHost, cfg.ServerPort)

	db, err := database.NewDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	log.Println("Connected to database")

	// Testing registration
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)

	log.Println("Services initialized")

	ctx := context.Background()
	user, err := authService.Register(ctx, "test@example.com", "password123")
	if err != nil {
		log.Fatalf("Registration failed: %v", err)
	}

	log.Printf("User registerd successfully: %+v", user)

	user, err = authService.Login(ctx, "test@example.com", "password123")
	if err != nil {
		log.Fatal("Login failed: %v", err)
	}

	log.Printf("User logged in successfully: %+v", user)
}
