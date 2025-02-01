package routes

import (
	"github.com/RiteshDevOpsEngineer/ecom/internal/middleware"
	"github.com/RiteshDevOpsEngineer/ecom/internal/services/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoutes(router *gin.Engine, jwtSecret []byte, mongoClient *mongo.Client, redisClient *redis.Client) {
	api := router.Group("/auth")
	{
		api.POST("/logout", middleware.AuthMiddleware(redisClient, mongoClient), auth.Logout(redisClient, mongoClient))
		api.POST("/login", auth.Login(mongoClient, jwtSecret, "github.com/RiteshDevOpsEngineer/ecom", "users"))
		api.POST("/verify-otp", auth.OtpVerify(redisClient, jwtSecret, mongoClient, "github.com/RiteshDevOpsEngineer/ecom", "users"))
	}
}
