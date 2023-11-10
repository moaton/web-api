package service

import (
	"net/http"

	"github.com/moaton/web-api/pkg/logger"
)

type Middleware interface {
	GetTokenFromHeader(r *http.Request) (string, error)
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (m *middleware) GetTokenFromHeader(r *http.Request) (string, error) {
	token, err := r.Cookie("Bearer")
	if err != nil {
		logger.Errorf("GetTokenFromHeader err %v", err)
		return "", err
	}
	email, err := Decode(token.Value)
	return email, err
}
