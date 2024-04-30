# Tron wallet v2

![image](https://github.com/ranjbar-dev/tron-wallet/blob/main/assets/image.png?raw=true)


V2 is all about simplicity! I'm revamping the package to be as easy to use as possible.


### crypto.go 

This file contains functions for working with cryptographic keys. It includes functions for generating private keys, public keys, and converting between them. These keys are likely used for signing transactions or other cryptographic operations in a blockchain or cryptocurrency application.

`GeneratePrivateKey() (*ecdsa.PrivateKey, error)`

`PrivateKeyToHex(privateKey *ecdsa.PrivateKey) string`

`PrivateKeyFromHex(privateKeyHex string) (*ecdsa.PrivateKey, error)`

`PrivateKeyToPublicKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error)`

`PublicKeyToHex(publicKey *ecdsa.PublicKey) string`

### address.go

This file contains functions for generating cryptographic addresses, likely for use in a blockchain or cryptocurrency application.

`PublicKeyToAddressHex(publicKey *ecdsa.PublicKey) string`

`PublicKeyToAddressBase58(publicKey *ecdsa.PublicKey) string`

