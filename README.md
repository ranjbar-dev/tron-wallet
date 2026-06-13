# Tron wallet v2

![image](https://github.com/ranjbar-dev/tron-wallet/blob/main/assets/image.png?raw=true)

V2 is all about simplicity! I'm revamping the package to be as easy to use as possible.

### Install

```
go get github.com/ranjbar-dev/tron-wallet/v2@latest
```

```go
import tronwallet "github.com/ranjbar-dev/tron-wallet/v2"
```

> v2 uses the `/v2` module path (Go semantic import versioning). The v1 series stays available at `github.com/ranjbar-dev/tron-wallet`.

### private_key.go 

Avaiable functions related to private key 

```

GeneratePrivateKey() (*ecdsa.PrivateKey, error)

PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string

PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) []byte

PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error)

PrivateKeyFromBytes(privateKeyBytes []byte) (*ecdsa.PrivateKey, error)

PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error)

```

### address.go

Avaiable functions related to address 

```

PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string

PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string

```

### grpc_client.go

Connect to a TRON node. `NewGrpcClient` returns a **started** client with the optional TronGrid API key already applied, so it is ready to pass straight to the transaction helpers. Remember to `defer client.Stop()`.

```

NewGrpcClient(grpcAddress string, timeout time.Duration, apiToken string) (*client.GrpcClient, error)

```

### transaction.go

Avaiable functions related to transaction 

you can create any transaction from `github.com/fbsobreira/gotron-sdk` and sign and broadcast it using `SignTransaction` and `BroadcastTransaction` 

```

CreateTransferTransaction(client *client.GrpcClient, from, to string, amount *big.Int) (*api.TransactionExtention, error)

CreateTRC20TransferTransaction(client *client.GrpcClient, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error)

CreateFreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error)

CreateUnfreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error)

SignTransaction(transaction *api.TransactionExtention, privateKey *ecdsa.PrivateKey) (*api.TransactionExtention, error)

BroadcastTransaction(client *client.GrpcClient, transaction *api.TransactionExtention) (*api.Return, error)

```

`CreateFreezeTransaction` / `CreateUnfreezeTransaction` use the TRON **Stake 2.0** endpoints (`FreezeBalanceV2` / `UnfreezeBalanceV2`). Stake 1.0 was disabled on mainnet, so the legacy freeze returns `freeze v2 is open, old freeze is closed` — these helpers avoid that.

### Amounts and precision

Every amount is expressed in the asset's smallest unit and is passed as a `*big.Int`:

- **TRC20** transfers support **arbitrary precision** — the amount is forwarded to the contract as a `*big.Int`, so 18-decimal tokens and values larger than `int64` work as expected.
- **TRX** transfers and **staking** are bounded by the TRON protocol to the `int64` range (the whole TRX supply fits comfortably), so those amounts are converted to `int64`. The conversion is guarded: a `nil`, negative, or out-of-range amount returns a clear error instead of silently overflowing.

### Fee explanation

- TRX transfer: if account has enough bandwidth to covert trasaction fee it is free( 5000 points ), if not it will cost 0.1 TRX. 

- TRC20 transfer: consume Energy and Bandwidth, if insufficient, TRX is burned as fee.

```
Fee ≈ Energy Used × Energy Unit Price (in TRX)
```


### Example 

Send 1 TRX on the Nile testnet — create, sign and broadcast:

```go
package main

import (
	"log"
	"math/big"
	"time"

	tronwallet "github.com/ranjbar-dev/tron-wallet/v2"
)

func main() {
	// Connect (apiToken is optional; pass "" to run without one).
	client, err := tronwallet.NewGrpcClient("grpc.nile.trongrid.io:50051", 10*time.Second, "")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Stop()

	// Load the sender's key.
	privateKey, err := tronwallet.PrivateKeyFromHex("YOUR_PRIVATE_KEY_HEX")
	if err != nil {
		log.Fatal(err)
	}

	from := "TQm6MTevKxyyKXzudM6UGjYqxnUmx2HiY3"
	to := "TEkxPcAR7GtkTvr8uQFgUsaFenFE2djkHB"

	// 1) create — amount is in sun (1 TRX = 1,000,000 sun).
	tx, err := tronwallet.CreateTransferTransaction(client, from, to, big.NewInt(1_000_000))
	if err != nil {
		log.Fatal(err)
	}

	// 2) sign
	tx, err = tronwallet.SignTransaction(tx, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 3) broadcast
	if _, err := tronwallet.BroadcastTransaction(client, tx); err != nil {
		log.Fatal(err)
	}
}
```

See the test files for more examples and usage of this package.

### Tests

The test suite is split into two layers:

- **Unit tests** are deterministic and run fully offline:

  ```
  go test ./...
  ```

- **Integration tests** talk to the live Nile testnet over gRPC and are guarded by a build tag, so they never run by accident:

  ```
  go test -tags=integration ./...
  ```

### TRX Faucet

Follow TronTestnet Twitter account @TronTest2 . Write your address in your tweet and @TronTest2 . They will transfer 10,000 test TRX (usually within five minutes). Each address can only be obtained once a day. If you need TRX for the nile testnet, please add "NILE" in your tweet.

### TRC20 Faucet

Go to https://developers.tron.network/ and connect to the discord community. You can than ask for usdt in #faucet channel. Just type !shasta_usdt YOUR_WALLET_ADDRESS and send. TronFAQ bot will send you 5000 USDT (SASHTA) soon.

### Important

I simplified this repository https://github.com/fbsobreira repository to create this package You can check go tron sdk for better examples and functionalities.

### TODOS 

- estimating trc20 and trc10 fee 

- contract api calls, deploy, get name, decimals and ... 

- block and transaction crawl calls 

- delegate resource transaction 


### License

Released under the [MIT License](LICENSE).

### Donation

Address `TUE66D1BT79FQcJE7fwf5vdfu3BYM4ph9Y`