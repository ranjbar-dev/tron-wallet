package tronWallet

import (
	"crypto/ecdsa"

	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/ranjbar-dev/tron-wallet/grpcClient"
)

func estimateTrc10TransactionFee(node enums.Node, privateKey *ecdsa.PrivateKey, fromAddressBase58 string, toAddressBase58 string, amountInSun int64) (int64, error) {

	tx, err := createTransactionInput(node, fromAddressBase58, toAddressBase58, amountInSun)
	if err != nil {
		return 0, err
	}

	singedTx, err := signTransaction(tx, privateKey)
	if err != nil {
		return 0, err
	}

	temp := (len(singedTx.Transaction.Signature[0]) + len(singedTx.Transaction.RawData.String())) / 2
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

func estimateTrc20TransactionFee() (int64, error) {

	return trc20FeeLimit, nil
}
