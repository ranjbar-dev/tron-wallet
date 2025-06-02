package tronwallet

import (
	"math/big"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCreateTransferTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	_, err = CreateTransferTransaction(conn, fromAddress, toAddress, big.NewInt(1000000))
	if err != nil {

		t.Fatalf("CreateTransferTransaction failed: %v", err)
	}

}

func TestCreateTRC20TransferTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	_, err = CreateTRC20TransferTransaction(conn, fromAddress, toAddress, contractAddress, big.NewInt(1000000), 1000000)
	if err != nil {

		t.Fatalf("CreateTRC20TransferTransaction failed: %v", err)
	}

}

func TestCreateFreezTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	_, err = CreateFreezTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000))
	if err != nil {

		t.Fatalf("CreateFreezTransaction failed: %v", err)
	}

}

func TestCreateUnfreezeTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	_, err = CreateUnfreezeTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000))
	if err != nil {

		t.Fatalf("CreateUnfreezeTransaction failed: %v", err)
	}

}

func TestSignTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	// Here can be any type of transaction transfer, freez, unfreeze, etc.
	transaction, err := CreateFreezTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000))
	if err != nil {

		t.Fatalf("CreateFreezTransaction failed: %v", err)
	}

	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	transaction, err = SignTransaction(transaction, privateKey)
	if err != nil {

		t.Fatalf("SignTransaction failed: %v", err)
	}
}

func TestBroadcastTransaction(t *testing.T) {

	conn := client.NewGrpcClient("grpc.nile.trongrid.io:50051")
	err := conn.Start(grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		t.Fatalf("Failed to start gRPC client: %v", err)
	}
	defer conn.Stop()

	transaction, err := CreateFreezTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000))
	if err != nil {

		t.Fatalf("CreateFreezTransaction failed: %v", err)
	}

	privateKey, err := PrivateKeyFromHex(fromAddressPrivateKeyHex)
	if err != nil {

		t.Fatalf("PrivateKeyFromHex failed: %v", err)
	}

	transaction, err = SignTransaction(transaction, privateKey)
	if err != nil {

		t.Fatalf("SignTransaction failed: %v", err)
	}

	_, err = BroadcastTransaction(conn, transaction)
	if err != nil {

		t.Fatalf("BroadcastTransaction failed: %v", err)
	}

}
