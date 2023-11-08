package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/moaton/web-api/pkg/logger"
)

var secret string = "MYSECRETKEY"

func Encode(id int64, email string) (map[string]string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   id,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})
	at, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		logger.Errorf("NewToken accessToken err %v", err)
		return map[string]string{}, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		logger.Errorf("NewToken refreshToken err %v", err)
		return map[string]string{}, err
	}

	return map[string]string{
		"access_token":  at,
		"refresh_token": rt,
	}, nil
}

func Decode(token string) (string, error) {
	tokenType, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		logger.Errorf("Decode ParseWithClaims err %v", err)
		return "", err
	}

	claims := tokenType.Claims.(jwt.MapClaims)
	return claims["email"].(string), nil
}
