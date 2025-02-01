package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"erspl/internal/adapters/repository/redis"
	"erspl/internal/core/domain"
	"erspl/internal/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(mongoClient *mongo.Client, jwtSecret []byte, dbName, userCollection string) gin.HandlerFunc {

	return func(c *gin.Context) {
		var request struct {
			Mobile string `json:"mobile" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			utils.NewResponse(http.StatusBadRequest, "false", "Invalid request", nil).Send(c)
			return
		}

		mobile := request.Mobile
		if !utils.IsValidPhone(mobile) {
			utils.NewResponse(http.StatusBadRequest, "error", "Invalid mobile number", nil).Send(c)
			return
		}

		collection := mongoClient.Database(dbName).Collection(userCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var user domain.User
		err := collection.FindOne(ctx, bson.M{"phone": mobile}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				user = domain.User{
					Phone:       mobile,
					Status:      1,
					OTPCount:    0,
					LastOTPSent: 0,
				}
				_, err = collection.InsertOne(ctx, user)
				if err != nil {
					utils.NewResponse(http.StatusInternalServerError, "false", "Database error", nil).Send(c)
					return
				}
			} else {
				utils.NewResponse(http.StatusInternalServerError, "false", "Database error", nil).Send(c)
				return
			}
		}

		if user.Status == 2 {
			utils.NewResponse(http.StatusUnauthorized, "error", "Your account is suspended", nil).Send(c)
			return
		}

		if !utils.CanSendOTP(&user) {
			utils.NewResponse(http.StatusTooManyRequests, "false", "Too many requests", nil).Send(c)
			return
		}

		utils.UpdateOTPSent(&user)
		user.OTP = utils.GenerateOTP()
		_, err = collection.UpdateOne(ctx, bson.M{"phone": mobile}, bson.M{"$set": user})
		if err != nil {
			utils.NewResponse(http.StatusInternalServerError, "false", "Database error", nil).Send(c)
			return
		}

		redisClient := redis.NewGoRedisClient()
		redisClient.Set(context.Background(), mobile, user.OTP, domain.OTPExpireTime)

		message := fmt.Sprintf("Welcome to eRSPL! Use OTP %s to verify your account. Regards, eRSPL", user.OTP)
		err = utils.SendSMS(mobile, message)
		if err != nil {
			utils.NewResponse(http.StatusInternalServerError, "false", "Failed to send OTP", nil).Send(c)
			return
		}

		utils.NewResponse(http.StatusOK, "true", "OTP sent", nil).Send(c)
	}
}
