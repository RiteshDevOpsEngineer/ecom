package middleware

import (
	"context"
	"erspl/config"
	"erspl/internal/core/domain"
	"erspl/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthMiddleware(redisClient *redis.Client, mongoClient *mongo.Client) gin.HandlerFunc {
	cfg, _ := config.New()

	jwtSecret := []byte(cfg.JWTSecretKey)
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.NewResponse(http.StatusUnauthorized, "false", "Authorization header required", nil).Send(c)
			c.Abort()
			return
		}

		claims := &domain.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.NewResponse(http.StatusUnauthorized, "false", "Invalid token", nil).Send(c)
			c.Abort()
			return
		}

		// Check if the token exists in Redis
		_, err = redisClient.Get(context.Background(), claims.Phone).Result()
		if err != nil {
			dbName := "erspl"
			CollectionName := "users"
			userCollection := mongoClient.Database(dbName).Collection(CollectionName)

			// Define a context with a timeout
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Query the user collection to find the user with the matching phone number and token
			var user domain.User
			err := userCollection.FindOne(ctx, bson.M{"phone": claims.Phone, "token": tokenString}).Decode(&user)
			if err != nil {
				// If an error occurs or the user is not found, send an error response
				utils.NewResponse(http.StatusUnauthorized, "false", "Token expired or not found", nil).Send(c)
				c.Abort()
				return
			}

			c.Set("phone", claims.Phone)
			c.Next()
			return
		}

		c.Set("phone", claims.Phone)
		c.Next()
	}
}

// // internal/middleware/auth.go
// package middleware

// import (
// 	"context"
// 	"erspl/config"
// 	"erspl/internal/core/domain"
// 	"erspl/internal/utils"
// 	"fmt"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8"
// )

// func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
// 	cfg, _ := config.New()

// 	jwtSecret := []byte(cfg.JWTSecretKey)
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Authorization header required", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		claims := &domain.Claims{}
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			// Check the signing method and secret key
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return jwtSecret, nil
// 		})

// 		if err != nil || !token.Valid {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Invalid token", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		_, err = redisClient.Get(context.Background(), claims.Phone).Result()
// 		if err != nil {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Token expired or not found", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		c.Set("phone", claims.Phone)
// 		c.Next()
// 	}
// }

// func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc {
// 	cfg, _ := config.New()

// 	jwtSecret := []byte(cfg.JWTSecretKey)
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Authorization header required", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		claims := &domain.Claims{}
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtSecret, nil
// 		})

// 		if err != nil || !token.Valid {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Invalid token", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		_, err = redisClient.Get(context.Background(), claims.Phone).Result()
// 		if err != nil {
// 			utils.NewResponse(http.StatusUnauthorized, "false", "Token expired or not found", nil).Send(c)
// 			c.Abort()
// 			return
// 		}

// 		c.Set("phone", claims.Phone)
// 		c.Next()
// 	}
// }
