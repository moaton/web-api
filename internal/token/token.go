package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/moaton/web-api/internal/models"
	"github.com/moaton/web-api/pkg/logger"
)

type Token interface {
	CreateAccessToken(id int64, email string) (string, error)
	CreateRefreshToken(id int64) (string, error)
	IsAuthorized(token string) (bool, error)
	ExtractIDFromToken(requestToken string) (int64, error)
}

type token struct {
	secret               string
	accessTokenExpMinute int64
	refreshTokenExpDays  int64
}

func New(secret string, accessTokenExpMinute, refreshTokenExpDays int64) Token {
	return &token{
		secret:               secret,
		accessTokenExpMinute: accessTokenExpMinute,
		refreshTokenExpDays:  refreshTokenExpDays,
	}
}

func (t *token) CreateAccessToken(id int64, email string) (string, error) {
	exp := time.Now().Add(time.Duration(t.accessTokenExpMinute) * time.Minute)
	claims := models.JwtCustomClaims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (t *token) CreateRefreshToken(id int64) (string, error) {
	exp := time.Now().Add(time.Duration(t.accessTokenExpMinute) * time.Minute)
	claimsRefresh := models.JwtCustomClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	refreshToken, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (t *token) IsAuthorized(token string) (bool, error) {
	_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(jt *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		logger.Errorf("Token IsAuthorized ParseWithClaims err %v", err)
		return false, err
	}

	return true, nil
}

func (t *token) ExtractIDFromToken(requestToken string) (int64, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return 0, fmt.Errorf("invalid Token")
	}
	return claims["id"].(int64), nil
}
