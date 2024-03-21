package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	// Setup a Gin router
	router := gin.Default()
	router.POST("/login", LoginHandler)

	// Prepare a test user
	testUser := `{"name": "testUser", "password": "testPassword"}`

	// Test case: Valid login request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(testUser))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Test case: Invalid login request (user doesn't exist)
	invalidUser := `{"name": "nonExistentUser", "password": "testPassword"}`
	req, _ = http.NewRequest("POST", "/login", bytes.NewBufferString(invalidUser))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)

	// Add more test cases as needed...
}
