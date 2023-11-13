package middleware_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moaton/web-api/internal/middleware"
	"github.com/moaton/web-api/mocks"
	"github.com/moaton/web-api/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO
func TestAuthMiddleware(t *testing.T) {
	t1 := mocks.NewToken(t)
	mw := middleware.New("123", t1)

	//	Case: 401 Unauthorized (without token)
	req := httptest.NewRequest(http.MethodGet, "/revenue", nil)
	w := httptest.NewRecorder()

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseOk(w, "success")
	})

	handlerTest := mw.AuthMiddleware(nextHandler)

	handlerTest.ServeHTTP(w, req)
	result := w.Result()

	var response struct {
		Error string `json:"error"`
	}
	decoder := json.NewDecoder(result.Body)
	err := decoder.Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, 401, result.StatusCode, "they should be equal")
	assert.Equal(t, "token not found", response.Error, "they should be equal")
	result.Body.Close()

	//	Case: 401 Unauthorized (with token) | IsAuthorized error
	t1.On("CreateAccessToken", int64(5), "test").Return("test", nil)
	t1.On("IsAuthorized", "test").Return(false, errors.New("error"))

	accessToken, err := t1.CreateAccessToken(5, "test")
	require.NoError(t, err)

	fmt.Println(accessToken)

	req = httptest.NewRequest(http.MethodGet, "/revenue", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	w = httptest.NewRecorder()

	handlerTest = mw.AuthMiddleware(nextHandler)

	handlerTest.ServeHTTP(w, req)
	result = w.Result()

	decoder = json.NewDecoder(result.Body)
	err = decoder.Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, 401, result.StatusCode, "they should be equal")
	assert.Equal(t, "error", response.Error, "they should be equal")
	result.Body.Close()

	//	Case: 401 Unauthorized (with token) | IsAuthorized !ok
	t2 := mocks.NewToken(t)
	mw = middleware.New("123", t2)
	t2.On("CreateAccessToken", int64(5), "test").Return("test", nil)
	t2.On("IsAuthorized", "test").Return(false, nil)

	accessToken, err = t2.CreateAccessToken(5, "test")
	require.NoError(t, err)

	fmt.Println(accessToken)

	req = httptest.NewRequest(http.MethodGet, "/revenue", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	w = httptest.NewRecorder()

	handlerTest = mw.AuthMiddleware(nextHandler)

	handlerTest.ServeHTTP(w, req)
	result = w.Result()

	decoder = json.NewDecoder(result.Body)
	err = decoder.Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, 401, result.StatusCode, "they should be equal")
	assert.Equal(t, "", response.Error, "they should be equal")
	result.Body.Close()

	//	Case: Success
	t3 := mocks.NewToken(t)
	mw = middleware.New("123", t3)
	t3.On("CreateAccessToken", int64(5), "test").Return("test", nil)
	t3.On("IsAuthorized", "test").Return(true, nil)

	accessToken, err = t3.CreateAccessToken(5, "test")
	require.NoError(t, err)

	fmt.Println(accessToken)

	req = httptest.NewRequest(http.MethodGet, "/revenue", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	w = httptest.NewRecorder()

	handlerTest = mw.AuthMiddleware(nextHandler)

	handlerTest.ServeHTTP(w, req)
	result = w.Result()
	var res string
	decoder = json.NewDecoder(result.Body)
	err = decoder.Decode(&res)
	require.NoError(t, err)

	assert.Equal(t, 200, result.StatusCode, "they should be equal")
	assert.Equal(t, "success", res, "they should be equal")
	result.Body.Close()

}
