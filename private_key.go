package tronwallet

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GeneratePrivateKey
// Generate a new private key
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

// PrivateKeyToHex
// Convert private key to hex string
func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// PrivateKeyToBytes
// Convert private key to bytes
func PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) []byte {

	return crypto.FromECDSA(privateKey)
}

// PrivateKeyFromHex
// Convert hex string to private key
func PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(privateKeyHex)
}

// PrivateKeyFromBytes
// Convert bytes to private key
func PrivateKeyFromBytes(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {

	return crypto.ToECDSA(privateKeyBytes)
}

// PrivateKeyToPublicKey
// Get public key from private key
func PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error getting public key from private key")
	}

	return publicKeyECDSA, nil
}
