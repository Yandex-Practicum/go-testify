package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	expected := strings.Join(cafeList["moscow"], ",")
	assert.GreaterOrEqual(t, 10, totalCount)                  // Проверяет, что введенное число больше, чем есть в списке
	assert.Equal(t, http.StatusOK, responseRecorder.Code)     // Проверяет, что статус ответа 200 OK
	assert.Equal(t, expected, responseRecorder.Body.String()) // Проверяет, что вернулись все кафе
}
func TestMainHandlerWhenCityNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=taganrog", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)       // Проверяет, что статус ответа 400 BadRequest
	assert.Equal(t, "wrong city value", responseRecorder.Body.String()) // Проверяет сообщение об ошибке
}
func TestMainHandlerWhenReqIsValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.NotEmpty(t, responseRecorder.Body.String())    //  Убеждается, что тело ответа не пустое
	require.Equal(t, http.StatusOK, responseRecorder.Code) // Проверяет, что статус ответа 200 OK
	returnedCafe := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, returnedCafe, 3)
}
