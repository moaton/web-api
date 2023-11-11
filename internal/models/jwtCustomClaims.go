package models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID int64 `json:"id"`
	jwt.RegisteredClaims
}
