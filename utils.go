package tronwallet

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/mr-tron/base58"
)

// validateAmount rejects a nil or negative amount. It is the shared precondition
// for every amount that crosses the package boundary.
func validateAmount(amount *big.Int) error {
	if amount == nil {
		return errors.New("amount must not be nil")
	}
	if amount.Sign() < 0 {
		return fmt.Errorf("amount must not be negative, got %s", amount.String())
	}

	return nil
}

// amountToInt64 converts a TRON amount (in the asset's smallest unit) to the
// int64 the gRPC API expects for native TRX and staking operations. Those values
// are bounded by the protocol to the int64 range, so this guards against a nil,
// negative or out-of-range amount with a clear error instead of silently
// wrapping through big.Int.Int64.
//
// TRC20 amounts are not converted here: they are passed to the SDK as *big.Int
// and therefore support arbitrary precision.
func amountToInt64(amount *big.Int) (int64, error) {
	if err := validateAmount(amount); err != nil {
		return 0, err
	}
	if !amount.IsInt64() {
		return 0, fmt.Errorf("amount %s exceeds the int64 range supported by the TRON protocol", amount.String())
	}

	return amount.Int64(), nil
}

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
