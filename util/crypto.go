package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/mr-tron/base58"
)

func S256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	bs := h.Sum(nil)
	return bs
}

func HexToBase58(str string) string {

	addb, _ := hex.DecodeString(str)
	hash1 := S256(S256(addb))
	secret := hash1[:4]
	for _, v := range secret {
		addb = append(addb, v)
	}
	return base58.Encode(addb)
}
