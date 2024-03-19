package network

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

// SetupTest creates a random address for testing.
func SetupTest(t *testing.T) string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	port := random.Intn(1000) + 50000 // Random port between 50000 and 50999
	address := fmt.Sprintf("localhost:%d", port)

	return address
}

// TestMain is the main entry point for tests.
func TestMain(m *testing.M) {
	// Add any global setup code here if needed
	exitVal := m.Run()
	// Add any global teardown code here if needed
	os.Exit(exitVal)
}
