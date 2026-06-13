package tronwallet

import (
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
)

// NewGrpcClient builds a gotron-sdk gRPC client configured with a per-call
// timeout and an optional TronGrid API token.
//
// grpcAddress is the node endpoint, e.g. "grpc.nile.trongrid.io:50051".
// timeout bounds every RPC, e.g. 10 * time.Second.
// apiToken is a TronGrid API key; pass "" to run without one.
func NewGrpcClient(grpcAddress string, timeout time.Duration, apiToken string) *client.GrpcClient {
	c := client.NewGrpcClientWithTimeout(grpcAddress, timeout)

	_ = c.SetAPIKey(apiToken)

	return c
}
