// Package tronwallet is a small, dependency-light toolkit for working with the
// TRON network in Go. It wraps github.com/fbsobreira/gotron-sdk and exposes a
// deliberately minimal surface so that the common wallet tasks read top to
// bottom with no hidden state.
//
// # What it covers
//
//   - Keys: generate secp256k1 private keys and convert between hex, bytes and
//     public keys (privatekey.go).
//   - Addresses: derive hex and Base58Check ("T...") addresses from a public key
//     (address.go).
//   - Transactions: build, sign and broadcast TRX transfers, TRC20 transfers and
//     Stake 2.0 freeze/unfreeze transactions (transaction.go).
//   - Connectivity: construct a configured gRPC client, optionally with a
//     TronGrid API key (grpc_client.go).
//
// # Units
//
// Every amount is expressed in the asset's smallest unit. TRX is measured in sun
// (1 TRX = 1,000,000 sun). TRC20 amounts use the token's own decimals
// (amount = value * 10^decimals); because token amounts can exceed an int64, the
// TRC20 helper accepts a *big.Int and supports arbitrary precision.
//
// # Lifecycle
//
// The transaction helpers follow the same three steps:
//
//	create -> sign -> broadcast
//
// Building a transaction never moves funds; only a signed, broadcast transaction
// does. A typical end-to-end flow:
//
//	c, err := tronwallet.NewGrpcClient("grpc.nile.trongrid.io:50051", 10*time.Second, "")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer c.Stop()
//
//	pk, _ := tronwallet.PrivateKeyFromHex(senderHex)
//
//	tx, err := tronwallet.CreateTransferTransaction(c, from, to, big.NewInt(1_000_000)) // 1 TRX
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	tx, err = tronwallet.SignTransaction(tx, pk)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	if _, err = tronwallet.BroadcastTransaction(c, tx); err != nil {
//		log.Fatal(err)
//	}
//
// For functionality beyond this package (contract deployment, block crawling,
// resource delegation, ...) use github.com/fbsobreira/gotron-sdk directly.
package tronwallet
