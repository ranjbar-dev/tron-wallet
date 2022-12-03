package tronWallet

import (
	"crypto/ecdsa"
	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/ranjbar-dev/tron-wallet/grpcClient"
)

func EstimateTrc10TransactionFee(node enums.Node, privateKey *ecdsa.PrivateKey, fromAddressBase58 string, toAddressBase58 string, amountInSun int64) (int64, error) {

	tx, err := createTransactionInput(node, fromAddressBase58, toAddressBase58, amountInSun)

	tx, err = signTransaction(tx, privateKey)

	temp := (len(tx.Transaction.Signature[0]) + len(tx.Transaction.RawData.String())) / 2
	bandwidthNeed := int64(temp + 68)

	c, _ := grpcClient.GetGrpcClient(enums.SHASTA_NODE)

	res, err := c.GetAccountResource(fromAddressBase58)
	if err != nil {
		return 0, err
	}

	avaialable := res.FreeNetLimit - res.FreeNetUsed

	if avaialable > bandwidthNeed {
		return 0, nil
	}

	return bandwidthNeed * 1000, err
}

func EstimateTrc20TransactionFee(node enums.Node, privateKey *ecdsa.PrivateKey, fromAddressBase58 string, toAddressBase58 string, amountInSun int64) (int64, error) {

	return 0, nil
}
