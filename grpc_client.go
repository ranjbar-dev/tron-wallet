package tronwallet

import (
	"fmt"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGrpcClient dials grpcAddress and returns a started, ready-to-use gRPC
// client. The returned client already has the TronGrid API token applied (when
// provided) and an open connection, so it can be passed straight to the
// transaction helpers. Close it with client.Stop when you are done.
//
// grpcAddress is the node endpoint, e.g. "grpc.nile.trongrid.io:50051".
// timeout bounds every RPC, e.g. 10 * time.Second.
// apiToken is a TronGrid API key sent as the TRON-PRO-API-KEY header on every
// call; pass "" to run without one.
//
// The connection is opened with insecure (plaintext) transport credentials,
// which is what the public TronGrid gRPC endpoints expect on port 50051. To talk
// to a TLS endpoint, build the client manually with client.NewGrpcClient and
// call Start with your own grpc.DialOption values.
func NewGrpcClient(grpcAddress string, timeout time.Duration, apiToken string) (*client.GrpcClient, error) {
	c := client.NewGrpcClientWithTimeout(grpcAddress, timeout)

	if apiToken != "" {
		if err := c.SetAPIKey(apiToken); err != nil {
			return nil, fmt.Errorf("set api key: %w", err)
		}
	}

	if err := c.Start(grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, fmt.Errorf("start grpc client: %w", err)
	}

	return c, nil
}
