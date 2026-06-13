package tronwallet

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// PublicKeyToAddressHex derives the TRON address for a public key and returns it
// as a 42-character lowercase hex string. A TRON address is the Keccak-256 hash
// of the public key (the last 20 bytes) prefixed with the network byte 0x41,
// which is why the result always starts with "41".
func PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string {
	address := crypto.PubkeyToAddress(*publicKey).Hex()

	address = "41" + address[2:]

	return strings.ToLower(address)
}

// PublicKeyToAddressBase58 derives the TRON address for a public key and returns
// it in the Base58Check form used throughout the TRON ecosystem (the familiar
// "T..." address). This is the encoding you pass to the transaction helpers.
func PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string {
	return hexToBase58(PublicKeyToAddressHex(publicKey))
}
