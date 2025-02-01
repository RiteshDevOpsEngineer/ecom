package main

import (
	"erspl/config"
	"erspl/internal/adapters/database"
	"erspl/internal/middleware"
	"erspl/internal/routes"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, _ := config.New()

	jwtSecret := []byte(cfg.JWTSecretKey)

	mongoClient := database.GetMongoClient()
	redisClient := database.GetRedisClient()

	router := gin.Default()

	router.Use(middleware.MaintenanceMiddleware(cfg))

	routes.SetupAuthRoutes(router, jwtSecret, mongoClient, redisClient)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
