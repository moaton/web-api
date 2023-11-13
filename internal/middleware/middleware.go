package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/moaton/web-api/internal/token"

	"github.com/moaton/web-api/pkg/utils"
)

type Middleware interface {
	AuthMiddleware(next http.Handler) http.Handler
}

type middleware struct {
	secret string
	token  token.Token
}

func New(secret string, token token.Token) Middleware {
	return &middleware{
		secret: secret,
		token:  token,
	}
}

func (m *middleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		t := strings.Split(token, " ")
		if len(t) != 2 {
			utils.ResponseError(w, http.StatusUnauthorized, "token not found")
			return
		}

		ok, err := m.token.IsAuthorized(t[1])
		fmt.Println("err ", err)
		if err != nil {
			utils.ResponseError(w, http.StatusUnauthorized, err.Error())
			return
		}
		if !ok {
			utils.ResponseError(w, http.StatusUnauthorized, "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
