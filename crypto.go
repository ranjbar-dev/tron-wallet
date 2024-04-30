package tronwallet

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}

func hexToBase58(str string) string {

	addb, _ := hex.DecodeString(str)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}
	return base58.Encode(addb)
}

func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

func PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(privateKeyHex)
}

func PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error getting public key from private key")
	}

	return publicKeyECDSA, nil
}

func PublicKeyToHex(publicKey *ecdsa.PublicKey) string {

	privateKeyBytes := crypto.FromECDSAPub(publicKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}
