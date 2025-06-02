package tronwallet

import (
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {

	_, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}
}

func TestPrivateKeyToHex(t *testing.T) {

	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	hexStr := PrivateKeyToHex(privateKey)
	if len(hexStr) != 64 {

		t.Errorf("Expected hex string length of 64, got %d", len(hexStr))
	}
}

func TestPrivateKeyToBytes(t *testing.T) {

	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	bytes := PrivateKeyToBytes(privateKey)
	if len(bytes) != 32 {

		t.Errorf("Expected bytes length of 32, got %d", len(bytes))
	}
}

func TestPrivateKeyFromHex(t *testing.T) {

	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	hexStr := PrivateKeyToHex(privateKey)
	if len(hexStr) != 64 {

		t.Errorf("Expected hex string length of 64, got %d", len(hexStr))
	}

	_, err = PrivateKeyFromHex(hexStr)
	if err != nil {

		t.Fatalf("PrivateKeyFromHex failed: %v", err)
	}
}

func TestPrivateKeyFromBytes(t *testing.T) {

	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("GeneratePrivateKey failed: %v", err)
	}

	bytes := PrivateKeyToBytes(privateKey)
	if len(bytes) != 32 {

		t.Errorf("Expected bytes length of 32, got %d", len(bytes))
	}

	_, err = PrivateKeyFromBytes(bytes)
	if err != nil {

		t.Fatalf("PrivateKeyFromBytes failed: %v", err)
	}
}

func TestPrivateKeyToPublicKey(t *testing.T) {

	privKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("Failed to generate private key: %v", err)
	}

	_, err = PrivateKeyToPublicKey(privKey)
	if err != nil {

		t.Fatalf("PrivateKeyToPublicKey failed: %v", err)
	}
}
