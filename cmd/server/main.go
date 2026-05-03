package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/umarov-safar/user-management-api/internal/config"
	"github.com/umarov-safar/user-management-api/internal/database"
	"github.com/umarov-safar/user-management-api/internal/routes"

	_ "github.com/lib/pq"
)

func main() {
	// loading config
	cfg := config.LoadConfig()
	log.Printf("%+v \n", cfg)
	log.Printf("Starting server on %s:%s", cfg.ServerHost, cfg.ServerPort)

	// connecting to db
	db, err := database.NewDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	log.Println("Connected to database")

	// routes registration
	router := gin.Default()

	routes.RegisterRoutes(router, db, cfg)

	log.Println("Routes registerd")

	// running server
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
