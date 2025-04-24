package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверяем правильность кода ответа
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	//проверяем что тело ответа не пустое
	assert.NotEmpty(t, responseRecorder.Body)
}
func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?count=2&city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	//проверяем правильность кода ответа
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	bodyExpected := "wrong city value"
	//проверяем правильность возвращаемой ошибки в теле ответа
	assert.Equal(t, bodyExpected, responseRecorder.Body.String())

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//проверяем, что количество кафе, равно общему количеству кафе
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)
}
