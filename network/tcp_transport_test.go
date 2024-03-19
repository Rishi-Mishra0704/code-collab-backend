package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport_Listen(t *testing.T) {
	transport := NewTCPTransport()
	address := SetupTest(t)

	// Test successful listener initialization
	err := transport.Listen(address)
	assert.NoError(t, err)
	assert.NotNil(t, transport.Listener)

	// Test attempting to listen on the same address again
	err = transport.Listen(address)
	assert.NoError(t, err) // No error should occur since listener is already started
}

func TestTCPTransport_Close(t *testing.T) {
	transport := NewTCPTransport()
	address := SetupTest(t)

	// Initialize listener
	err := transport.Listen(address)
	assert.NoError(t, err)

	// Test closing the listener
	err = transport.Close()
	assert.NoError(t, err)
	assert.Nil(t, transport.Listener)
}

func TestTCPTransport_Close_NotInitialized(t *testing.T) {
	transport := NewTCPTransport()

	// Test closing when listener is not initialized
	err := transport.Close()
	assert.NoError(t, err) // No error should occur if listener is not initialized
}

func TestTCPTransport_Listen_Error(t *testing.T) {
	transport := NewTCPTransport()
	invalidAddress := "invalid_address"

	// Test listener initialization with an invalid address
	err := transport.Listen(invalidAddress)
	assert.Error(t, err)
	assert.Nil(t, transport.Listener)
}
