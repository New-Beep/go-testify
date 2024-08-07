package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenOK(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	statusOK := http.StatusOK
	status := responseRecorder.Code

	require.NotEmpty(t, responseRecorder.Body.String())
	if !assert.Equal(t, status, statusOK) {
		t.Errorf("expected status code: %d, got %d", statusOK, status)
	}
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedAnswer := `wrong city value`
	answer, err := io.ReadAll(responseRecorder.Body)
	if err != nil {
		t.Error(err)
	}
	status := http.StatusBadRequest
	city := req.URL.Query().Get("city")
	if !assert.Equal(t, city, "moscow") {
		assert.Equal(t, status, responseRecorder.Code)
		assert.Equal(t, expectedAnswer, string(answer))
	}
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	countStr := req.URL.Query().Get("count")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	require.NotEqual(t, countStr, "")
	assert.GreaterOrEqual(t, len(list), totalCount)

}
