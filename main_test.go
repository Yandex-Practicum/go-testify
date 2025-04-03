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

	responseRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRec, req)

	require.Equal(t, http.StatusOK, responseRec.Code)
	assert.NotEmpty(t, responseRec.Body.String())
	assert.Len(t, strings.Split(responseRec.Body.String(), ","), totalCount)
}
func TestMainHandlerRequestValid(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRec, req)

	require.Equal(t, http.StatusOK, responseRec.Code)
	assert.NotEmpty(t, responseRec.Body.String())
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unknowncity", nil)

	responseRec := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRec, req)

	require.Equal(t, http.StatusBadRequest, responseRec.Code)
	assert.Equal(t, "wrong city value", responseRec.Body.String())
}
