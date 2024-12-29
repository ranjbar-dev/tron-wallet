package tronwallet

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string {

	address := crypto.PubkeyToAddress(*publicKey).Hex()

	address = "41" + address[2:]

	return strings.ToLower(address)
}

func PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string {

	return hexToBase58(PublicKeyToAddressHex(publicKey))
}
