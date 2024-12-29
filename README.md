# Tron wallet v2

![image](https://github.com/ranjbar-dev/tron-wallet/blob/main/assets/image.png?raw=true)


V2 is all about simplicity! I'm revamping the package to be as easy to use as possible.

### private_key.go 

Avaiable functions related to private key 

`GeneratePrivateKey() (*ecdsa.PrivateKey, error)`

`PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string`

`PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) []byte`

`PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error)`

`PrivateKeyFromBytes(privateKeyBytes []byte) (*ecdsa.PrivateKey, error)`

`PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error)`

### address.go

Avaiable functions related to address 

`PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string`

`PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string`


### transaction.go

Avaiable functions related to transaction 

you can create any transaction from `github.com/fbsobreira` and sign and broadcast it using `SignTransaction` and `BroadcastTransaction` 

`CreateTransferTransaction(client *client.GrpcClient, from, to string, amount *big.Int) (*api.TransactionExtention, error)`

`CreateTRC20TransferTransaction(client *client.GrpcClient, from, to, contract string, amount *big.Int, feeLimit int64) (*api.TransactionExtention, error)`

`CreateFreezTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error)`

`CreateUnfreezeTransaction(client *client.GrpcClient, address string, resource core.ResourceCode, amount *big.Int) (*api.TransactionExtention, error)`

`SignTransaction(transaction *api.TransactionExtention, privateKey *ecdsa.PrivateKey) (*api.TransactionExtention, error)`

`BroadcastTransaction(client *client.GrpcClient, transaction *api.TransactionExtention) (*api.Return, error)`

### TODOS 

- estimating trc20 and trc10 fee 

- contract api calls, deploy, get name, decimals and ... 

- block and transaction crawl calls 

- delegate resource transaction 

- writing test 

### TRX Faucet

Follow TronTestnet Twitter account @TronTest2 . Write your address in your tweet and @TronTest2 . They will transfer 10,000 test TRX (usually within five minutes). Each address can only be obtained once a day. If you need TRX for the nile testnet, please add "NILE" in your tweet.

### TRC20 Faucet

Go to https://developers.tron.network/ and connect to the discord community. You can than ask for usdt in #faucet channel. Just type !shasta_usdt YOUR_WALLET_ADDRESS and send. TronFAQ bot will send you 5000 USDT (SASHTA) soon.

### Important

I simplified this repository https://github.com/fbsobreira repository to create this package You can check go tron sdk for better examples and functionalities.

### Donation

Address `TUE66D1BT79FQcJE7fwf5vdfu3BYM4ph9Y`