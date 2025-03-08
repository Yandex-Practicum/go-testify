package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	body := responseRecorder.Body.String()

	require.NotEmpty(t, body)

	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusOK)

	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
}

func TestMainHandlerWhenUnsupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=kazan", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// необходимые проверки
	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	body := responseRecorder.Body.String()

	expectedBody := "wrong city value"
	assert.Equal(t, body, expectedBody)
}
