package tronwallet

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/mr-tron/base58"
)

// s256 returns the SHA-256 digest of s.
func s256(s []byte) []byte {
	h := sha256.New()
	h.Write(s)

	return h.Sum(nil)
}

// hexToBase58 converts a hex-encoded TRON address (e.g. "41..." ) into its
// Base58Check representation. The 4-byte checksum is the first four bytes of the
// double SHA-256 of the address, appended before Base58 encoding.
func hexToBase58(str string) string {
	addb, _ := hex.DecodeString(str)
	checksum := s256(s256(addb))[:4]
	addb = append(addb, checksum...)

	return base58.Encode(addb)
}
