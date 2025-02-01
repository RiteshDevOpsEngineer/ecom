package routes

import (
	"erspl/internal/middleware"
	"erspl/internal/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoutes(router *gin.Engine, jwtSecret []byte, mongoClient *mongo.Client, redisClient *redis.Client) {
	api := router.Group("/auth")
	{
		api.POST("/logout", middleware.AuthMiddleware(redisClient, mongoClient), auth.Logout(redisClient, mongoClient))
		api.POST("/login", auth.Login(mongoClient, jwtSecret, "erspl", "users"))
		api.POST("/verify-otp", auth.OtpVerify(redisClient, jwtSecret, mongoClient, "erspl", "users"))
	}
}
