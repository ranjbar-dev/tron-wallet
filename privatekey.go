package tronwallet

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GeneratePrivateKey generates a new random secp256k1 private key suitable for a
// TRON account. The returned key never touches the network; derive its address
// with PrivateKeyToPublicKey followed by PublicKeyToAddressBase58.
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

// PrivateKeyToHex encodes a private key as a 64-character lowercase hex string
// without the "0x" prefix.
func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// PrivateKeyToBytes returns the raw 32-byte big-endian encoding of a private key.
func PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) []byte {
	return crypto.FromECDSA(privateKey)
}

// PrivateKeyFromHex parses a hex-encoded private key (with or without the "0x"
// prefix) back into an *ecdsa.PrivateKey.
func PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKeyHex)
}

// PrivateKeyFromBytes parses the raw 32-byte big-endian encoding of a private
// key back into an *ecdsa.PrivateKey.
func PrivateKeyFromBytes(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {
	return crypto.ToECDSA(privateKeyBytes)
}

// PrivateKeyToPublicKey returns the public key that corresponds to privateKey.
func PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error getting public key from private key")
	}

	return publicKeyECDSA, nil
}
