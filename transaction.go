package tronwallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func CreateTrxTransferTransaction(fromAddress Address, toAddress Address, sun int64) *Transaction {

	// TODO : implement

	temp := Transaction{
		RawData: RawData{
			Contract: []Contract{
				{
					Type: "TransferContract", // TODO : use const for these fields later
					Parameter: ContractParameter{
						TypeUrl: "type.googleapis.com/protocol.TransferContract", // TODO : use const for these fields later
						Value: ContractValue{
							OwnerAddress: fromAddress.Base58,
							ToAddress:    toAddress.Hex,
							Amount:       sun,
						},
					},
				},
			},
			RefBlockBytes: "",
			RefBlockHash:  "",
			Expiration:    0,
			Timestamp:     0,
		},
	}

	return &temp
}

func SignTransactionHex(hex []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {

	signature, err := crypto.Sign(hex, privateKey)
	if err != nil {
		return nil, fmt.Errorf("sign error: %v", err)
	}

	return signature, nil
}

func BroadcastSingedTransaction(signedTransaction []byte) error {

	// TODO : implement

	return nil
}
