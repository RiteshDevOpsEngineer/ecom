package domain

type User struct {
	Phone       string `bson:"phone" json:"phone"`
	Status      int    `bson:"status" json:"status"`
	OTP         string `bson:"otp" json:"otp"`
	OTPCount    int64  `bson:"otp_count" json:"otp_count"`
	LastOTPSent int64  `bson:"last_otp_sent" json:"last_otp_sent"`
	Token       string `bson:"token" json:"token"`
}
