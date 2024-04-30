package tronwallet

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func PublicKeyToAddress(publicKey *ecdsa.PublicKey) *Address {

	temp := Address{
		Hex:    PublicKeyToAddressHex(publicKey),
		Base58: PublicKeyToAddressBase58(publicKey),
	}

	return &temp
}

func PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string {

	address := crypto.PubkeyToAddress(*publicKey).Hex()

	address = "41" + address[2:]

	return strings.ToLower(address)
}

func PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string {

	addressHex := PublicKeyToAddressHex(publicKey)

	return hexToBase58(addressHex)
}
