package tronwallet

import (
	"strings"
	"testing"
)

func TestPublicKeyToAddressHex(t *testing.T) {

	// Generate a test private key
	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Get the public key
	pubKey, err := PrivateKeyToPublicKey(privateKey)
	if err != nil {

		t.Fatalf("Failed to generate public key: %v", err)
	}

	// Get the address
	address := PublicKeyToAddressHex(pubKey)

	// Basic validation checks
	if len(address) != 42 {

		t.Errorf("Expected address length of 42, got %d", len(address))
	}

	if address[:2] != "41" {

		t.Errorf("Expected address to start with '41', got %s", address[:2])
	}
}

func TestPublicKeyToAddressBase58(t *testing.T) {

	// Generate a test private key
	privateKey, err := GeneratePrivateKey()
	if err != nil {

		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Get the public key
	pubKey, err := PrivateKeyToPublicKey(privateKey)
	if err != nil {

		t.Fatalf("Failed to generate public key: %v", err)
	}

	// Get the base58 address
	address := PublicKeyToAddressBase58(pubKey)

	// Basic validation checks
	if len(address) == 0 {

		t.Error("Expected non-empty address")
	}

	// Base58 addresses should not contain characters like 0, O, I, l
	invalidChars := "0OIl"
	for _, char := range invalidChars {

		if strings.Contains(address, string(char)) {

			t.Errorf("Address contains invalid Base58 character: %c", char)
		}
	}
}
