package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignupHandler(t *testing.T) {
	// Setup a Gin router
	router := gin.Default()
	router.POST("/signup", SignupHandler)

	// Test case: Valid signup request
	validPayload := `{"id": "user1", "name": "John Doe", "email": "john@example.com", "password": "password123"}`
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBufferString(validPayload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Test case: Invalid request payload
	invalidPayload := `{"id": "user1", "email": "john@example.com", "password": "password123"}`
	req, _ = http.NewRequest("POST", "/signup", bytes.NewBufferString(invalidPayload))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

}
