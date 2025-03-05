package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func TestMainHandler_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")
	assert.NotEmpty(t, responseRecorder.Body.String(), "Ответ не должен быть пустым")
}

func TestMainHandler_WrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=paris", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался статус 400")
	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "Неверное сообщение об ошибке")
}

func TestMainHandler_CountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedCafeList := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался статус 200")
	assert.Equal(t, expectedCafeList, responseRecorder.Body.String(), "Ожидался полный список кафе")
}

func main() {
	http.HandleFunc("/cafe", mainHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



