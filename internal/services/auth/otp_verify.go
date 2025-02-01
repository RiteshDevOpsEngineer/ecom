package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/RiteshDevOpsEngineer/ecom/internal/core/domain"
	"github.com/RiteshDevOpsEngineer/ecom/internal/utils"
)

func OtpVerify(redisClient *redis.Client, jwtSecret []byte, mongoClient *mongo.Client, dbName, userCollection string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Mobile string `json:"mobile"`
			OTP    string `json:"otp"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			utils.NewResponse(http.StatusBadRequest, "false", "Invalid request", nil).Send(c)
			return
		}

		otp, err := redisClient.Get(context.Background(), request.Mobile).Result()
		if err != nil || otp != request.OTP {
			utils.NewResponse(http.StatusUnauthorized, "false", "Invalid OTP", nil).Send(c)
			return
		}

		expirationTime := time.Now().Add(domain.JWTExpireTime)
		claims := &domain.Claims{
			Phone: request.Mobile,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			utils.NewResponse(http.StatusInternalServerError, "false", "Could not generate token", nil).Send(c)
			return
		}

		err = redisClient.Set(context.Background(), request.Mobile, tokenString, domain.JWTExpireTime).Err()
		if err != nil {
			utils.NewResponse(http.StatusInternalServerError, "false", "Failed to store token", nil).Send(c)
			return
		}

		collection := mongoClient.Database(dbName).Collection(userCollection)
		filter := bson.M{"phone": request.Mobile}
		update := bson.M{"$set": bson.M{"token": tokenString}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			utils.NewResponse(http.StatusInternalServerError, "false", "Failed to store token", nil).Send(c)
			return
		}

		utils.NewResponse(http.StatusOK, "true", "Token generated successfully", gin.H{"token": tokenString}).Send(c)
	}
}
