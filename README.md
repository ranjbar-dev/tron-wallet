# tron-wallet
tron wallet package for creating and generating wallet, transferring TRX, getting wallet balance and crawling blocks to find wallet transactions

### Installation 
```
go get github.com/ranjbar-dev/tron-wallet@v1.0.2
```

### Test
test for `Crawler`, `TronWallet` and `Token` located at `/test`
```
go test ./test
```

### Wallet methods 
- generating tron wallet 
```
w := GenerateTronWallet(enums.SHASTA_NODE)
w.Address // strnig 
w.AddressBase58 // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```
- creating tron wallet from private key 
```
w := CreateTronWallet(enums.SHASTA_NODE,privateKeyHex)
w.Address // strnig 
w.AddressBase58 // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
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