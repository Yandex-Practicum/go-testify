package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Invalid status code")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Response body is empty")
}

func TestMainHandlerInvalidCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=vorkuta&count=1", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Invalid status code")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Unexpected response body")
}

func TestMainHandlerCountMoreThanMax(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", fmt.Sprintf("/cafe?count=%d&city=moscow", totalCount), nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Invalid status code")
	assert.Equal(t, "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент", responseRecorder.Body.String(), "Unexpected response body")
}
