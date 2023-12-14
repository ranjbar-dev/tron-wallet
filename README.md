![image](https://github.com/ranjbar-dev/tron-wallet/blob/main/assets/image.png?raw=true)


# tron-wallet
tron wallet package for creating and generating wallet, transferring TRX, getting wallet balance and crawling blocks to find wallet transactions

### Installation 
```
go get github.com/ranjbar-dev/tron-wallet@v1.0.6
```

### Test
test for `Crawler`, `TronWallet` and `Token` located at `/test`
```
go test ./test
```

### Wallet methods 
**Set [environment variables](https://developers.tron.network/reference/select-network) `TRON_PRO_API_KEY` in production if you want to avoid rate limit.**

- generating tron wallet 
```go
w := GenerateTronWallet(enums.SHASTA_NODE)
w.Address // string 
w.AddressBase58 // string 
w.PrivateKey // string 
w.PublicKey // string 
```
- creating tron wallet from private key 
```go
w := CreateTronWallet(enums.SHASTA_NODE,privateKeyHex)
w.Address // string 
w.AddressBase58 // string 
w.PrivateKey // string 
w.PublicKey // string 
```

- generating mnemonic 
```go
mnenomic := tronWallet.GenerateMnemonic(12)
// net uncle rigid useless coast explain saddle crawl pupil erase veteran slender
```
- creating tron wallet from mnemonic with account path
```go
mnemonic := "net uncle rigid useless coast explain saddle crawl pupil erase veteran slender"
wallet, _ := tronWallet.MnemonicToTronWallet(enums.NILE_NODE, mnemonic, "m/44'/195'/3'/0/1", "")

wallet.AddressBase58 // TXTaWVTCMAEjC35S6sLF5gi6ZKVrxAkmGX
wallet.Address  // 41ebb83dedb47dc852a5e2863acaf7b11989bc07a9
wallet.PrivateKey // 900b8fc4c8c83a9baffc40917aa1a029eb4b75215d05d0de92e365b907f27c22
wallet.PublicKey //04487ff8ed9de594a4148dfe0f83b7320e069fa66848f078f90270b695022c671af47417004b4cdd53487e8def2ebb6fe696fd883e48d68a0ed1bed9a3459f4a01
```

- getting wallet trx balance 
```
balanceInSun,err := w.Balance()
balanceInSun // int64 
```
- getting wallet trc20 balance
```
balanceInToken,err := w.BalanceTRC20(token)
balanceInToken // int64 
```
- crawl blocks for addresses transactions 
```

c := &Crawler{
		Node: enums.SHASTA_NODE, // network -> maninet, shasta, nile
		Addresses: []string{
			"TY3PJu3VY8xVUc5BjYwJtyRgP7TfivV666", // list of your addresses
		},
	}
	
res, err := c.ScanBlocks(40) // scan latest 40 block on block chain and extract addressess transactions 
res, err := c.ScanBlocksFromTo(28905305, 28905307) // or scan block from to number
Example 
// *
{
    {
        "address": "TY3PJu3VY8xVUc5BjYwJtyRgP7TfivV666",
        "tranasctions": {
            {
                "tx_id": "6afbc5758d49e8d8bedddd903edbfc01c5f11ebfbaa6237e887294a6fc9394a2",
                "from_address": "TJnsY5bGiwuPCQFismQDwyVTPAn7M88888",
                "to_address": "TY3PJu3VY8xVUc5BjYwJtyRgP7TfivV666",
                "amount": "195500", // sun
                "confirmations": 1,
                "symbol": "TRX",
            },
            {
                "tx_id": "61624fffb31d09b9fe7252cd743733a5890c2a9077698d4a9bcb3d70ebb28777",
                "from_address": "TSw5FSuWhTAcaJmBUVFY9fUY4ihwx588b6",
                "to_address": "TJnsY5bGiwuPCQFismQDwyVTPAn7M88888",
                "amount": "10000", // in token sub value 
                "confirmations": 1,
                "symbol": "USDT",
            },
        }
    },
    ...
}
* // 
	
```

Estimate transfer trx fee
```
feeInSun,err := w.EstimateTransferFee("TJnsY5bGiwuPCQFismQDwyVTPAn7M88888",10000)
feeInSun // int64
```

Estimate transfer trc20 fee
```
feeInSun,err := w.EstimateTransferTRC20Fee()
feeInSun // int64
```

- transfer trx from wallet 
```
txId, err := w.Transfer(toAddressBase58, amount)
txId // string 
```
- transfer trc20 from wallet
```
txId, err := w.TransferTRC20(token, toAddressBase58, amount)
txId // string 
```

### Token methods 
- declaring token 
```
token := &tronWallet.Token{
    ContractAddress: enums.SHASTA_Tether_USDT,
}
```
- Getting token name 
```
token.GetName(w.Node, w.AddressBase58) // return string,error
``` 
- Getting token symbol
```
token.GetSymbol(w.Node, w.AddressBase58) // return string,error
```
- Getting token decimals 
```
token.GetDecimal(w.Node, w.AddressBase58) // return int64,error
```

### Supported networks
check `enums/nodes` file
alternatively you can create your own node
```
node := enums.CreateNode("grpc.test.com:50051")
```

### Supported contracts
check `enums/contracts` file
alternatively you can create your own contract
```
contractAddress := enums.CreateContractAddress("TPYmHEhy5n8TCEfYGqW2rPxsghSfzghPDn")
```

### TRX Faucet
Follow TronTestnet Twitter account
@TronTest2
.
Write your address in your tweet and
@TronTest2
.
They will transfer 10,000 test TRX (usually within five minutes).
Each address can only be obtained once a day.
If you need TRX for the nile testnet, please add "NILE" in your tweet.

### TRC20 Faucet
Go to https://developers.tron.network/ and connect to the discord community.
You can than ask for usdt in #faucet channel.
Just type !shasta_usdt YOUR_WALLET_ADDRESS and send. TronFAQ bot will send you 5000  USDT (SASHTA) soon.


### Important
I simplified this repository https://github.com/fbsobreira repository to create this package
You can check go tron sdk for better examples and functionalities
and do not use this package in production, I created this package for education purposes.


### Donation
Address `TUE66D1BT79FQcJE7fwf5vdfu3BYM4ph9Y`
