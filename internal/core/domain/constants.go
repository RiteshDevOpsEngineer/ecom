package domain

import "time"

const (
	MaxOTPPer10Minutes = 5
	OTPExpireTime      = 5 * time.Minute
	JWTExpireTime      = 1 * time.Hour
)
