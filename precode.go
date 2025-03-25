package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=2", nil)
	rr := httptest.NewRecorder()
	mainHandle(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "Код ответа не 200")
	assert.NotEmpty(t, rr.Body.String(), "Тело ответа пустое")
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=london&count=2", nil)
	rr := httptest.NewRecorder()
	mainHandle(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code, "Код ответа не 400")
	assert.Equal(t, "wrong city value", rr.Body.String(), "Ошибка в теле ответа")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?city=moscow&count=10", nil)
	rr := httptest.NewRecorder()
	mainHandle(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "Код ответа не 200")

	cafes := strings.Split(rr.Body.String(), ",")
	assert.Len(t, cafes, 4, "Количество кафе не равно 4")
}