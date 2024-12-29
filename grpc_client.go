package tronwallet

import (
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
)

// NewGrpcClient
// Create a new grpc client using github.com/fbsobreira/gotron-sdk lib with timeout and api token
// grpcAddress is the address of the grpc server, example: grpc.nile.trongrid.io:50051
// timeout is the timeout for the grpc client, example: 10 * time.Second
// apiToken is the api token for the grpc client, you can leave it blank, example: ""
func NewGrpcClient(grpcAddress string, timeout time.Duration, apiToken string) *client.GrpcClient {

	c := client.NewGrpcClientWithTimeout(grpcAddress, timeout)

	_ = c.SetAPIKey(apiToken)

	return c
}
