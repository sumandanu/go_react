package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    return router
}

func TestHealthCheck(t *testing.T) {

	mockResponse := `{"status":"ok"}`

	router := SetUpRouter()
    router.GET("/health", healthCheck)

	req, _ := http.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    responseData, _ := io.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}