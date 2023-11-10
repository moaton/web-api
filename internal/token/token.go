package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/moaton/web-api/pkg/logger"
)

type Token interface {
	CreateAccessToken()
	CreateRefreshToken()
	IsAuthorized(token string) (bool, error)
}

type token struct {
	secret string
}

func New(secret string) Token {
	return &token{
		secret: secret,
	}
}

func (t *token) CreateAccessToken() {

}

func (t *token) CreateRefreshToken() {

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
