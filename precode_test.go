package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testHandler = http.HandlerFunc(mainHandle)

// strRequest создаем запрос, затем передаем в функцию executeRequest,
// чтобы получить объект ResponseRecorder.
func strRequest(method, url string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, url, body)
	return executeRequest(req)
}

// executeRequest выполняем обработку HTTP-запроса с помощью testHandler.
// Возвращаем записанный ответ.
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	testHandler.ServeHTTP(recorder, req)
	return recorder
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestHandler_Count(t *testing.T) {
	expectedCount := len(cafeList["moscow"])
	countCafe := expectedCount + 2

	response := strRequest("GET", fmt.Sprintf("/cafe?count=%d&city=moscow", countCafe), http.NoBody)

	require.Equal(t, http.StatusOK, response.Code)
	cafes := strings.Split(response.Body.String(), ",")
	assert.Len(t, cafes, expectedCount)
}

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestHandler_Success(t *testing.T) {
	response := strRequest("GET", "/cafe?count=2&city=moscow", http.NoBody)

	require.Equal(t, http.StatusOK, response.Code)
	assert.NotEmpty(t, response.Body)
}

// Город, который передаётся в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestHandler_InvalidCity(t *testing.T) {
	response := strRequest("GET", "/cafe?count=2&city=berlin", http.NoBody)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Contains(t, response.Body.String(), "wrong city value")
}
