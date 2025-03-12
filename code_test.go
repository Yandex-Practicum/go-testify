package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandleSuccessfulResponse(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEqual(t, responseRecorder.Body.Len(), 0)
}

// Город, который передаётся в параметре `city`, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку `wrong city value` в теле ответа.
func TestMainHandleIncorrectCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=hjkt", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	responseBody := "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), responseBody)
}

// Если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandleWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=200&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	responseSlice := strings.Split(responseRecorder.Body.String(), ",")
	assert.Equal(t, len(responseSlice), totalCount)
}
