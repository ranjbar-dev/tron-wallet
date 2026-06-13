//go:build integration

package tronwallet

import (
	"testing"
	"time"
)

// TestNewGrpcClient verifies that NewGrpcClient returns a fully initialized,
// connected client (issue #11): the connection is started and an RPC succeeds
// without the caller having to call Start or wire the API key by hand.
func TestNewGrpcClient(t *testing.T) {
	conn, err := NewGrpcClient(nileGrpcEndpoint, 10*time.Second, "")
	if err != nil {
		t.Fatalf("NewGrpcClient failed: %v", err)
	}
	defer conn.Stop()

	if _, err := conn.GetNowBlock(); err != nil {
		t.Fatalf("GetNowBlock failed: %v", err)
	}
}
