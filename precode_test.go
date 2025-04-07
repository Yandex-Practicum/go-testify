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

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count="+strconv.Itoa(totalCount)+"&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

	cafes := strings.Split(body, ",")
	assert.Len(t, cafes, 4)
}
