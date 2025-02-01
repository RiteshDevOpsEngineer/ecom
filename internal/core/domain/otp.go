package domain

import (
	"time"

	"gorm.io/gorm"
)

type OTP struct {
	gorm.Model
	UserID    uint
	Code      string
	ExpiresAt time.Time
}
