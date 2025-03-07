package main

import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
    "github.com/stretchr/testify/assert"
)

var cafeList = map[string][]string{
    "moscow": {"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
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

// Тесты
func TestValidRequestReturns200AndNotEmptyBody(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
    assert.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    assert.NotEmpty(t, responseRecorder.Body.String())
}

func TestInvalidCityReturns400(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=unknown&count=2", nil)
    assert.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
    assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestCountMoreThanTotalReturnsAllCafes(t *testing.T) {
    req, err := http.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
    assert.NoError(t, err)

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(responseRecorder, req)

    assert.Equal(t, http.StatusOK, responseRecorder.Code)
    expectedResponse := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
    assert.Equal(t, expectedResponse, responseRecorder.Body.String())
}

// Основная функция, которая запускает сервер
func main() {
    http.HandleFunc("/cafe", mainHandle)
    http.ListenAndServe(":8080", nil)
}


