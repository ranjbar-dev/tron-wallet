package tronWallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/ranjbar-dev/tron-wallet/grpcClient"
	"github.com/ranjbar-dev/tron-wallet/util"
	"math/big"
)

type TronWallet struct {
	Node          enums.Node
	Address       string
	AddressBase58 string
	PrivateKey    string
	PublicKey     string
}

// generating

func GenerateTronWallet(node enums.Node) *TronWallet {

	privateKey, _ := generatePrivateKey()
	privateKeyHex := convertPrivateKeyToHex(privateKey)

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)
	addressBase58 := util.HexToBase58(address)

	return &TronWallet{
		Node:          node,
		Address:       address,
		AddressBase58: addressBase58,
		PrivateKey:    privateKeyHex,
		PublicKey:     publicKeyHex,
	}
}

func CreateTronWallet(node enums.Node, privateKeyHex string) (*TronWallet, error) {

	privateKey, err := privateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)
	addressBase58 := util.HexToBase58(address)

	return &TronWallet{
		Node:          node,
		Address:       address,
		AddressBase58: addressBase58,
		PrivateKey:    privateKeyHex,
		PublicKey:     publicKeyHex,
	}, nil
}

// struct functions

func (t *TronWallet) PrivateKeyRCDSA() (*ecdsa.PrivateKey, error) {
	return privateKeyFromHex(t.PrivateKey)
}

func (t *TronWallet) PrivateKeyBytes() ([]byte, error) {

	priv, err := t.PrivateKeyRCDSA()
	if err != nil {
		return []byte{}, err
	}

	return crypto.FromECDSA(priv), nil
}

// private key

func generatePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

func convertPrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

func privateKeyFromHex(hex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(hex)
}

// public key

func getPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error in getting public key")
	}

	return publicKeyECDSA, nil
}

func convertPublicKeyToHex(publicKey *ecdsa.PublicKey) string {

	privateKeyBytes := crypto.FromECDSAPub(publicKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// address

func getAddressFromPublicKey(publicKey *ecdsa.PublicKey) string {

	address := crypto.PubkeyToAddress(*publicKey).Hex()

	address = "41" + address[2:]

	return address
}

// balance

func (t *TronWallet) Balance() (int64, error) {

	c, err := grpcClient.GetGrpcClient(t.Node)
	if err != nil {
		return 0, err
	}

	b, err := c.GetAccount(t.AddressBase58)
	if err != nil {
		return 0, err
	}

	return b.Balance, nil
}

func (t *TronWallet) BalanceTRC20(token *Token) (int64, error) {

	balance, err := token.GetBalance(t.Node, t.AddressBase58)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

// transaction

func (t *TronWallet) Transfer(toAddressBase58 string, amountInSun int64) (string, error) {

	privateRCDSA, err := t.PrivateKeyRCDSA()
	if err != nil {
		return "", fmt.Errorf("RCDSA private key error: %v", err)
	}

	tx, err := createTransactionInput(t.Node, t.AddressBase58, toAddressBase58, amountInSun)
	if err != nil {
		return "", fmt.Errorf("creating tx pb error: %v", err)
	}

	tx, err = signTransaction(tx, privateRCDSA)
	if err != nil {
		return "", fmt.Errorf("signing transaction error: %v", err)
	}

	err = broadcastTransaction(t.Node, tx)
	if err != nil {
		return "", fmt.Errorf("broadcast transaction error: %v", err)
	}

	return hexutil.Encode(tx.GetTxid())[2:], nil
}

func (t *TronWallet) TransferTRC20(token *Token, toAddressBase58 string, amountInTRC20 int64) (string, error) {

	privateKey, err := t.PrivateKeyRCDSA()
	if err != nil {
		return "", err
	}

	tx, err := createTrc20TransactionInput(t.Node, t.AddressBase58, token, toAddressBase58, big.NewInt(amountInTRC20))
	if err != nil {
		return "", err
	}

	signedTx, err := signTransaction(tx, privateKey)
	if err != nil {
		return "", err
	}

	err = broadcastTransaction(t.Node, signedTx)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(tx.GetTxid())[2:], nil
}
