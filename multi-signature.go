package tronWallet

import (
	"crypto/ecdsa"

	"github.com/ranjbar-dev/tron-wallet/enums"
)

func CreateAndBroadcastMultiTransaction(node enums.Node, fromAddressBase58 string, toAddressBase58 string, amountInSun int64, privateKeys []*ecdsa.PrivateKey, _ ecdsa.PrivateKey) error {

	transaction, err := createTransactionInput(node, fromAddressBase58, toAddressBase58, amountInSun)
	if err != nil {
		return err
	}

	for _, privateKey := range privateKeys {
		transaction, err = signTransaction(transaction, privateKey)
		if err != nil {
			return err
		}
	}

	return broadcastTransaction(node, transaction)
}
