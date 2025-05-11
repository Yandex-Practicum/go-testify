package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerOK проверяет корректный запрос: должны получить статус 200 и непустое тело.
func TestMainHandlerOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "ожидается код ответа 200")
	assert.NotEmpty(t, rr.Body.String(), "тело ответа не должно быть пустым")
}

// TestMainHandlerWrongCity проверяет ошибку для неподдерживаемого города.
// Ожидаем статус 400 и сообщение "wrong city value" в теле ответа.
func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=1&city=spb", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusBadRequest, rr.Code, "ожидается код ответа 400")
	assert.Equal(t, "wrong city value", strings.TrimSpace(rr.Body.String()))
}

// TestMainHandlerWhenCountMoreThanTotal проверяет, что при запросе количества кафе,
// превышающем общее число, возвращаются все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4 // согласно условию задачи в Москве всего 4 кафе
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code, "ожидается код ответа 200")

	cafes := strings.Split(rr.Body.String(), ",")
	assert.Len(t, cafes, totalCount, "должны вернуться все доступные кафе")
}
