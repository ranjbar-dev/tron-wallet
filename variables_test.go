//go:build integration

package tronwallet

import (
	"testing"
	"time"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
)

// nileGrpcEndpoint is the public TRON Nile testnet gRPC node used by the
// integration tests.
const nileGrpcEndpoint = "grpc.nile.trongrid.io:50051"

// Integration-test fixtures. The private key below is a disposable Nile
// testnet-only key used purely to exercise the signing/broadcast flow; it holds
// no mainnet value. Run these tests with: go test -tags=integration ./...
var (
	fromAddress              = "TQm6MTevKxyyKXzudM6UGjYqxnUmx2HiY3"
	fromAddressPrivateKeyHex = "39252406ac52ae99a289694c4f0f9284ac5a777b5ed4d35d0f45c5e99700a61a"
	toAddress                = "TEkxPcAR7GtkTvr8uQFgUsaFenFE2djkHB"
	contractAddress          = "TU2T8vpHZhCNY8fXGVaHyeZrKm8s6HEXWe"
)

// dialNileTestnet returns a started client connected to the Nile testnet via the
// package's own NewGrpcClient, and registers cleanup. It fails the test if the
// connection cannot start.
func dialNileTestnet(t *testing.T) *client.GrpcClient {
	t.Helper()

	conn, err := NewGrpcClient(nileGrpcEndpoint, 10*time.Second, "")
	if err != nil {
		t.Fatalf("failed to start gRPC client: %v", err)
	}
	t.Cleanup(conn.Stop)

	return conn
}
