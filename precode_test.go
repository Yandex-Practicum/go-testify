package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandler_ValidRequest(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city=moscow", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())
}

func TestMainHandler_InvalidCity(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city=spb", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Equal(t, "wrong city value", rec.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 10
	req := httptest.NewRequest(http.MethodGet, "/cafe?count="+strconv.Itoa(totalCount)+"&city=moscow", nil)
	rec := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rec, req)

	expected := strings.Join(cafeList["moscow"], ",")
	body := rec.Body.String()

	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, body)
	assert.Len(t, strings.Split(body, ","), len(cafeList["moscow"]))
}
