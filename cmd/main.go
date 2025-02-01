package main

import (
	"github.com/RiteshDevOpsEngineer/ecom/config"
	"github.com/RiteshDevOpsEngineer/ecom/internal/adapters/database"
	"github.com/RiteshDevOpsEngineer/ecom/internal/middleware"
	"github.com/RiteshDevOpsEngineer/ecom/internal/routes"

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
