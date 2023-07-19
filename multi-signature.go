package tronWallet

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/ranjbar-dev/tron-wallet/grpcClient"
)

func CreateAndBroadcastMultiTransaction(node enums.Node, fromAddressBase58 string, toAddressBase58 string, amountInSun int64, privateKeys []*ecdsa.PrivateKey, _ ecdsa.PrivateKey) (string, error) {

	transaction, err := createTransactionInput(node, fromAddressBase58, toAddressBase58, amountInSun)
	if err != nil {
		return "", err
	}

	for _, privateKey := range privateKeys {
		transaction, err = signTransaction(transaction, privateKey)
		if err != nil {
			return "", err
		}
	}

	c, err := grpcClient.GetGrpcClient(node)
	if err != nil {
		return "", err
	}

	res, err := c.Broadcast(transaction.Transaction)
	if err != nil {
		return "", err
	}

	if !res.Result {
		return "", errors.New(res.Code.String())
	}

	return string(res.Message), nil
}
