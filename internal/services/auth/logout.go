package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/RiteshDevOpsEngineer/ecom/internal/utils"
)

func Logout(redisClient *redis.Client, mongoClient *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		phone, exists := c.Get("phone")
		if !exists {
			utils.NewResponse(http.StatusBadRequest, "false", "Phone number not found in context", nil).Send(c)
			return
		}

		phoneStr, ok := phone.(string)
		if !ok {
			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to convert phone number to string", nil).Send(c)
			return
		}

		tokenString, err := redisClient.Get(context.Background(), phoneStr).Result()
		if err != nil && err != redis.Nil {
			fmt.Println("Error retrieving token from Redis:", err) // Log error
			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to retrieve token from Redis", nil).Send(c)
			return
		}

		if tokenString != "" {
			err = redisClient.Del(context.Background(), phoneStr).Err()
			if err != nil {
				fmt.Println("Error deleting token from Redis:", err) // Log error
				utils.NewResponse(http.StatusInternalServerError, "true", "Failed to logout", nil).Send(c)
				return
			}
		}

		dbName := "erspl"
		userCollection := "users"
		collection := mongoClient.Database(dbName).Collection(userCollection)
		updateResult, err := collection.UpdateOne(context.Background(), bson.M{"phone": phoneStr}, bson.M{"$unset": bson.M{"token": ""}})
		if err != nil {
			fmt.Println("Error updating token in MongoDB:", err) // Log error
			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to logout", nil).Send(c)
			return
		}

		if updateResult.MatchedCount == 0 {
			utils.NewResponse(http.StatusNotFound, "false", "Token not found in MongoDB", nil).Send(c)
			return
		}

		// Send successful logout response
		utils.NewResponse(http.StatusOK, "true", "Logged out successfully", nil).Send(c)
	}
}

// package auth

// import (
// 	"context"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"

// 	"erspl/internal/utils"
// )

// func Logout(redisClient *redis.Client, mongoClient *mongo.Client) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Retrieve the phone number from the request context
// 		phone, exists := c.Get("phone")
// 		if !exists {
// 			utils.NewResponse(http.StatusBadRequest, "false", "Phone number not found in context", nil).Send(c)
// 			return
// 		}

// 		// Convert phone number to string
// 		phoneStr, ok := phone.(string)
// 		if !ok {
// 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to convert phone number to string", nil).Send(c)
// 			return
// 		}

// 		// Check if the token exists in Redis
// 		tokenString, err := redisClient.Get(context.Background(), phoneStr).Result()
// 		if err != nil {
// 			fmt.Println("Error retrieving token from Redis:", err) // Log error
// 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to retrieve token from Redis", nil).Send(c)
// 			return
// 		}

// 		// If token exists in Redis, delete it
// 		if tokenString != "" {
// 			err = redisClient.Del(context.Background(), phoneStr).Err()
// 			if err != nil {
// 				fmt.Println("Error deleting token from Redis:", err) // Log error
// 				utils.NewResponse(http.StatusInternalServerError, "true", "Failed to logout", nil).Send(c)
// 				return
// 			}
// 		}

// 		dbName := "erspl"
// 		userCollection := "users"
// 		collection := mongoClient.Database(dbName).Collection(userCollection)
// 		_, err = collection.UpdateOne(context.Background(), bson.M{"phone": phoneStr}, bson.M{"$unset": bson.M{"token": ""}})
// 		if err != nil {
// 			fmt.Println("Error updating token in MongoDB:", err) // Log error
// 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to logout", nil).Send(c)
// 			return
// 		}

// 		// Send successful logout response
// 		utils.NewResponse(http.StatusOK, "true", "Logged out successfully", nil).Send(c)
// 	}
// }

// // func Logout(redisClient *redis.Client) gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		// Retrieve the phone number from the request context
// // 		phone, exists := c.Get("phone")
// // 		if !exists {
// // 			utils.NewResponse(http.StatusBadRequest, "false", "Phone number not found in context", nil).Send(c)
// // 			return
// // 		}

// // 		// Convert phone number to string
// // 		phoneStr, ok := phone.(string)
// // 		if !ok {
// // 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to convert phone number to string", nil).Send(c)
// // 			return
// // 		}

// // 		// Retrieve the token from Redis using the phone number as the key
// // 		tokenString, err := redisClient.Get(context.Background(), phoneStr).Result()
// // 		if err != nil {
// // 			fmt.Println("Error retrieving token from Redis:", err) // Log error
// // 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to retrieve token from Redis", nil).Send(c)
// // 			return
// // 		}

// // 		if tokenString == "" {
// // 			utils.NewResponse(http.StatusUnauthorized, "false", "Token not found in Redis", nil).Send(c)
// // 			return
// // 		}

// // 		// Delete the token (phone number) entry from Redis
// // 		err = redisClient.Del(context.Background(), phoneStr).Err()
// // 		if err != nil {
// // 			fmt.Println("Error deleting token from Redis:", err) // Log error
// // 			utils.NewResponse(http.StatusInternalServerError, "true", "Failed to logout", nil).Send(c)
// // 			return
// // 		}

// // 		// Send successful logout response
// // 		utils.NewResponse(http.StatusOK, "true", "Logged out successfully", nil).Send(c)
// // 	}
// // }
