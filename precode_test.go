package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе. В сервисе будет только один город moscow, в котором будет всего 4 кафе. Нужно реализовать три теста:

//Запрос сформирован корректно, сервис возвращает код ответа 200
//и тело ответа не пустое.

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	assert.NotEmpty(t, responseRecorder.Body)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
}

//Город, который передаётся в параметре city, не поддерживается.
//Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.

func TestMainHandlerWhenCityWrong(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=spb", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, "wrong city value", responseRecorder.Body.String())
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

}
