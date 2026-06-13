//go:build integration

package tronwallet

import (
	"math/big"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
)

func TestCreateTransferTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

	if _, err := CreateTransferTransaction(conn, fromAddress, toAddress, big.NewInt(1000000)); err != nil {
		t.Fatalf("CreateTransferTransaction failed: %v", err)
	}
}

func TestCreateTRC20TransferTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

	if _, err := CreateTRC20TransferTransaction(conn, fromAddress, toAddress, contractAddress, big.NewInt(1000000), 1000000); err != nil {
		t.Fatalf("CreateTRC20TransferTransaction failed: %v", err)
	}
}

func TestCreateFreezTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

	if _, err := CreateFreezTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000)); err != nil {
		t.Fatalf("CreateFreezTransaction failed: %v", err)
	}
}

func TestCreateUnfreezeTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

	if _, err := CreateUnfreezeTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000)); err != nil {
		t.Fatalf("CreateUnfreezeTransaction failed: %v", err)
	}
}

func TestSignTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

	// Any transaction type works here: transfer, freeze, unfreeze, etc.
	transaction, err := CreateFreezTransaction(conn, fromAddress, core.ResourceCode_ENERGY, big.NewInt(1000000))
	if err != nil {
		t.Fatalf("CreateFreezTransaction failed: %v", err)
	}

	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	if _, err = SignTransaction(transaction, privateKey); err != nil {
		t.Fatalf("SignTransaction failed: %v", err)
	}
}

func TestBroadcastTransaction(t *testing.T) {
	conn := dialNileTestnet(t)

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

	if _, err = BroadcastTransaction(conn, transaction); err != nil {
		t.Fatalf("BroadcastTransaction failed: %v", err)
	}
}
