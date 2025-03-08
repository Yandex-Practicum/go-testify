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
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code: %d, got: %d", http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String(), "Expected non empty body")
	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount, "Expected response length: %d, got: %d", totalCount, len(strings.Split(responseRecorder.Body.String(), ",")))
}

func TestMainHandlerWhenEverythingIsOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code: %d, got: %d", http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String(), "Expected non empty body")
}

func TestMainHandlerWhenCityIsWrong(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Expected status code: %d, got: %d", http.StatusBadRequest, responseRecorder.Code)
	require.Equal(t, "wrong city value", responseRecorder.Body.String(), "Expected response body: '%s', got: '%s'", "wrong city value", responseRecorder.Body.String())
}
