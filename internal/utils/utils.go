// internal/utils/utils.go
package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/RiteshDevOpsEngineer/ecom/internal/core/domain"
)

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}

// Response
func IsValidPhone(mobile string) bool {
	re := regexp.MustCompile(`^[6-9]\d{9}$`)
	return re.MatchString(mobile)
}

func CanSendOTP(user *domain.User) bool {
	now := time.Now().Unix()
	if now-user.LastOTPSent < 600 && user.OTPCount >= domain.MaxOTPPer10Minutes {
		return false
	}

	return user.Status == 1
}

func UpdateOTPSent(user *domain.User) {
	now := time.Now().Unix()
	if now-user.LastOTPSent > 600 {
		user.OTPCount = 0
	}
	user.OTPCount++
	user.LastOTPSent = now
}
