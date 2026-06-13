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

// CreateTransferTransaction builds an unsigned native TRX transfer. Sign it with
// SignTransaction and publish it with BroadcastTransaction.
//
// from and to are Base58 addresses (e.g. "TTTA2rCqLTDAEEb4UwPD34qLiJ6AUhgzRH").
// amount is denominated in sun (1 TRX = 1,000,000 sun).
func CreateTransferTransaction(client *client.GrpcClient, from, to string, amount *big.Int) (*api.TransactionExtention, error) {
	return client.Transfer(from, to, amount.Int64())
}

// CreateTRC20TransferTransaction builds an unsigned TRC20 token transfer. Sign it
// with SignTransaction and publish it with BroadcastTransaction.
//
// from and to are Base58 addresses; contract is the TRC20 contract address in
// Base58. amount is in the token's smallest unit (amount = value * 10^decimals).
// feeLimit is the maximum fee, in sun, the sender will pay for execution.
func CreateTRC20TransferTransaction(client *client.GrpcClient, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error) {
	return client.TRC20Send(from, to, contract, amount, feeLimit)
}

// CreateFreezeTransaction builds an unsigned stake (freeze) transaction that locks
// TRX to gain bandwidth or energy. Sign it with SignTransaction and publish it
// with BroadcastTransaction.
//
// It uses the TRON Stake 2.0 endpoint (FreezeBalanceV2). The legacy Stake 1.0
// freeze was disabled on mainnet, so building a freeze with the old endpoint
// fails with "Contract validate error : freeze v2 is open, old freeze is
// closed"; using the V2 endpoint here avoids that error.
//
// address is the staker's Base58 address. resource selects what the stake
// produces (core.ResourceCode_ENERGY or core.ResourceCode_BANDWIDTH). amount is
// the TRX to freeze, in sun.
func CreateFreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error) {
	return client.FreezeBalanceV2(address, resource, amount.Int64())
}

// CreateUnfreezeTransaction builds an unsigned unstake (unfreeze) transaction that
// releases previously frozen TRX. Sign it with SignTransaction and publish it
// with BroadcastTransaction.
//
// It uses the TRON Stake 2.0 endpoint (UnfreezeBalanceV2) and so pairs with
// CreateFreezeTransaction.
//
// address is the staker's Base58 address. resource must match the resource the
// stake was created for. amount is the TRX to unfreeze, in sun.
func CreateUnfreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error) {
	return client.UnfreezeBalanceV2(address, resource, amount.Int64())
}

// SignTransaction signs transaction in place with privateKey and returns it. Call
// it repeatedly with different keys to attach multiple signatures for multi-sig
// accounts.
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

// BroadcastTransaction submits a signed transaction to the network and returns the
// node's acknowledgement. A successful return means the transaction was accepted
// into the mempool, not that it has been confirmed in a block.
func BroadcastTransaction(client *client.GrpcClient, transaction *api.TransactionExtention) (*api.Return, error) {
	return client.Broadcast(transaction.Transaction)
}
