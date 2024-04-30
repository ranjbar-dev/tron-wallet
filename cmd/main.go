package main

import (
	"fmt"

	tronwallet "github.com/ranjbar-dev/tron-wallet"
)

func main() {

	privateKey, _ := tronwallet.GeneratePrivateKey()

	fmt.Println(privateKey)
}
