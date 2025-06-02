package tronwallet

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/mr-tron/base58"
)

// s256
// sha256 hash
func s256(s []byte) []byte {

	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}

// hexToBase58
// convert hex to base58
func hexToBase58(str string) string {

	addb, _ := hex.DecodeString(str)
	hash1 := s256(s256(addb))
	secret := hash1[:4]
	for _, v := range secret {

		addb = append(addb, v)
	}

	return base58.Encode(addb)
}
