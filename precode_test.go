package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	


	
)

func TestMainHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Неправильный код состояния")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Тело ответа не должно быть пустым")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=paris&count=2", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Неправильный код состояния")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Неожиданное тело ответа")
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	require.NoError(t, err, "Ошибка при создании запроса")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Неправильный код состояния")
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String(), "Неожиданное тело ответа")
}