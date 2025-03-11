package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=7&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	resLength := len(strings.Split(responseRecorder.Body.String(), ","))

	assert.Equal(t, totalCount, resLength)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=4&city=mos", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	expectedBody := "wrong city value"
	body := responseRecorder.Body.String()

	assert.Equal(t, expectedBody, body)

	// Либо проверить по корректному названию города, который ожидает сервер и переданного параметра в url
	// но тогда фиаско теста будет отображать его корректную работу)

	// expectedCity := "moscow"
	// city := req.URL.Query().Get("city")

	// assert.Equal(t, expectedCity, city)
}

func TestMainHandlerWhenRequestOK(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, body)

}
