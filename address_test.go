package tronwallet

import (
	"strings"
	"testing"
)

func TestPublicKeyToAddressHex(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	pubKey, err := PrivateKeyToPublicKey(privateKey)
	if err != nil {
		t.Fatalf("Failed to generate public key: %v", err)
	}

	address := PublicKeyToAddressHex(pubKey)

	if len(address) != 42 {
		t.Errorf("Expected address length of 42, got %d", len(address))
	}

	if address[:2] != "41" {
		t.Errorf("Expected address to start with '41', got %s", address[:2])
	}
}

func TestPublicKeyToAddressBase58(t *testing.T) {
	privateKey, err := GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	pubKey, err := PrivateKeyToPublicKey(privateKey)
	if err != nil {
		t.Fatalf("Failed to generate public key: %v", err)
	}

	address := PublicKeyToAddressBase58(pubKey)

	if len(address) == 0 {
		t.Error("Expected non-empty address")
	}

	// Base58 addresses never contain the ambiguous characters 0, O, I or l.
	invalidChars := "0OIl"
	for _, char := range invalidChars {
		if strings.Contains(address, string(char)) {
			t.Errorf("Address contains invalid Base58 character: %c", char)
		}
	}
}
