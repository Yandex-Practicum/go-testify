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
	returnCafe := strings.Split(responseRecorder.Body.String(), ",") // Разделяет тело ответа по ,
	require.Equal(t, http.StatusOK, responseRecorder.Code)           // Проверяет, что статус ответа 200 OK
	assert.Len(t, returnCafe, totalCount)                            // Проверят, что длинная, которая вернулась, соответствует ожидаемой
}
func TestMainHandlerWhenCityNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=taganrog", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)      // Проверяет, что статус ответа 400 BadRequest
	assert.Equal(t, "wrong city value", responseRecorder.Body.String()) // Проверяет сообщение об ошибке
}
func TestMainHandlerWhenReqIsValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusOK, responseRecorder.Code) // Проверяет, что статус ответа 200 OK
	assert.NotEmpty(t, responseRecorder.Body.String())     //  Убеждается, что тело ответа не пустое
	returnedCafe := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, returnedCafe, 3)
}
