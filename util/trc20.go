package util

import (
	"errors"
	"math/big"
	"strings"
)

const Trc20TransferMethodSignature = "a9059cbb"

type Trc20TokenTransfer struct {
	To    string
	Value big.Int
}

func ParseTrc20TokenTransfer(dataHex string) (Trc20TokenTransfer, bool) {

	isValid := isDataTrc20TokenTransfer(dataHex)
	if !isValid {
		return Trc20TokenTransfer{}, false
	}

	toAddressNonPaddedHex := dataHex[len(Trc20TransferMethodSignature) : 64+len(Trc20TransferMethodSignature)]

	toAddress, err := GainAddressFromPaddedHex(toAddressNonPaddedHex)

	if err != nil {
		return Trc20TokenTransfer{}, false
	}

	valueStr := dataHex[64+len(Trc20TransferMethodSignature):]
	value := new(big.Int)
	value.SetString(valueStr, 16)

	return Trc20TokenTransfer{
		To:    toAddress,
		Value: *value,
	}, true
}

func isDataTrc20TokenTransfer(dataHex string) bool {

	if len(dataHex) == 136 && strings.HasPrefix(dataHex, Trc20TransferMethodSignature) {
		return true
	}

	return false
}

func GainAddressFromPaddedHex(s string) (string, error) {

	if len(s) < 24 {
		return "", errors.New("invalid address")
	}
	subStr := s[24:]

	return HexToAddress("41" + subStr).String(), nil
}
