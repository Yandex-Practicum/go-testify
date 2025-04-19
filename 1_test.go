package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCorrect(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected code 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "response body should not be empty")
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=paris&count=2", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, responseRecorder.Body.String(), "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal2(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status 200 OK")
	cafes := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, cafes, totalCount, "expected that all the cafes should be returned")
}
