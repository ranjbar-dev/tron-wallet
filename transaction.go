package tronwallet

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"
)

// CreateTransferTransaction
// create trx transfer transcton you should sign it before broadcast
// client is the grpc client
// from is the sender address in base58, example: TTTA2rCqLTDAEEb4UwPD34qLiJ6AUhgzRH
// to is the receiver address in base58, example: TXNcHLpSMnSXPgt9yG1AWSo7iHuqb9rcJG
// amount is the amount of trx you want to transfer in sun, example: big.NewInt(1000000) // 1 TRX
func CreateTransferTransaction(client *client.GrpcClient, from, to string, amount *big.Int) (*api.TransactionExtention, error) {

	return client.Transfer(from, to, amount.Int64())
}

// CreateTRC20TransferTransaction
// create trc20 transfer transcton you should sign it before broadcast
// client is the grpc client
// from is the sender address in base58, example: TTTA2rCqLTDAEEb4UwPD34qLiJ6AUhgzRH
// to is the receiver address in base58, example: TXNcHLpSMnSXPgt9yG1AWSo7iHuqb9rcJG
// contract is the contract address in base58, example: TU2T8vpHZhCNY8fXGVaHyeZrKm8s6HEXWe
// amount is the amount of trc20 you want to transfer in smallest unit, example: big.NewInt(100000000) // if decimal is 8, it means 1 coin
func CreateTRC20TransferTransaction(client *client.GrpcClient, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error) {

	return client.TRC20Send(from, to, contract, amount, feeLimit)
}

// CreateFreezTransaction
// create freeze transaction you should sign it before broadcast
// client is the grpc client
// address is the address in base58, example: TTTA2rCqLTDAEEb4UwPD34qLiJ6AUhgzRH
// resource is the resource code, example: core.ResourceCode_ENERGY
// amount is the amount of resource you want to freeze in smallest unit, example: big.NewInt(1000000) // 1 TRX
func CreateFreezTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error) {

	return client.FreezeBalanceV2(address, resource, amount.Int64())
}

// CreateUnfreezeTransaction
// create unfreeze transaction you should sign it before broadcast
// client is the grpc client
// address is the address in base58, example: TTTA2rCqLTDAEEb4UwPD34qLiJ6AUhgzRH
// resource is the resource code, example: core.ResourceCode_ENERGY
// amount is the amount of resource you want to unfreeze in smallest unit, example: big.NewInt(1000000) // 1 TRX
func CreateUnfreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error) {

	return client.UnfreezeBalanceV2(address, resource, amount.Int64())
}

// SignTransaction
// sign the transaction with private key, you can sign transaction multiple times wih different private keys for multi-sig
// transaction is the transaction you want to sign created using CreateTransferTransaction or CreateTRC20TransferTransaction or other methods
// privateKey is the private key of the sender
func SignTransaction(transaction *api.TransactionExtention, privateKey *ecdsa.PrivateKey) (*api.TransactionExtention, error) {

	rawData, err := proto.Marshal(transaction.Transaction.GetRawData())
	if err != nil {

		return transaction, fmt.Errorf("proto marshal tx raw data error: %v", err)
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {

		return transaction, fmt.Errorf("sign error: %v", err)
	}

	transaction.Transaction.Signature = append(transaction.Transaction.Signature, signature)

	return transaction, nil
}

// BroadcastTransaction
// broadcast the signed transaction to the network
// client is the grpc client
// transaction is the signed transaction created using SignTransaction
func BroadcastTransaction(client *client.GrpcClient, transaction *api.TransactionExtention) (*api.Return, error) {

	return client.Broadcast(transaction.Transaction)
}
