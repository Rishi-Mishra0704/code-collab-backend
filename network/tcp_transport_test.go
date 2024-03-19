package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport_Listen(t *testing.T) {
	transport := NewTCPTransport()

	// Test successful listener initialization
	address := SetupTest(t) // Use random port
	err := transport.Listen(address)
	assert.NoError(t, err)
	assert.NotNil(t, transport.Listener)

	// Test attempting to listen on the same address again
	err = transport.Listen(address)
	assert.NoError(t, err) // No error should occur since listener is already started
}

func TestTCPTransport_Listen_Error(t *testing.T) {
	transport := NewTCPTransport()

	// Test listener initialization with an invalid address
	err := transport.Listen("invalid_address")
	assert.Error(t, err)
	assert.Nil(t, transport.Listener)
}

func TestTCPTransport_Close(t *testing.T) {
	transport := NewTCPTransport()
	addr := SetupTest(t)
	// Initialize listener
	err := transport.Listen(addr) // Use random port
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

func TestTCPTransport_Connect(t *testing.T) {
	transport := NewTCPTransport()
	address := SetupTest(t) // Use random available port for testing

	// Start the transport listener
	err := transport.Listen(address)
	assert.NoError(t, err)

	// Run the Connect method
	peer, err := transport.Connect(address)
	assert.NoError(t, err)

	// Check that the peer is added to the map of connected peers
	transport.Mutex.Lock()
	defer transport.Mutex.Unlock()
	assert.Contains(t, transport.Peers, address)

	// Check that the peer's connection is established
	assert.NotNil(t, peer.Conn)
	assert.True(t, peer.Online)
}

func TestTCPTransport_Connect_Error(t *testing.T) {
	transport := NewTCPTransport()
	address := "invalid_address"

	// Run the Connect method with an invalid address
	peer, err := transport.Connect(address)

	// Check that the error is not nil
	assert.Error(t, err)
	assert.Nil(t, peer)
}
