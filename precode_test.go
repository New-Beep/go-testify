package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	statusOK := http.StatusOK
	status := responseRecorder.Code
	countStr := req.URL.Query().Get("count")
	require.NotEmpty(t, responseRecorder.Body.String())
	require.NotEqual(t, countStr, "")
	assert.Equal(t, status, statusOK)

	city := req.URL.Query().Get("city")
	moscow := "moscow"
	if !assert.Equal(t, city, moscow) {
		fmt.Println("wrong city value")
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	if assert.GreaterOrEqual(t, len(list), totalCount) {
		fmt.Printf("В городе %s найдено %d заведения: %s", moscow, totalCount, list)
	}
}
