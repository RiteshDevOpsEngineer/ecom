package domain

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Phone string `json:"phone"`
	jwt.StandardClaims
}
