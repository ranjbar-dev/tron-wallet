package tronWallet

import (
	"fmt"
	"math/big"

	"github.com/ranjbar-dev/tron-wallet/enums"
	"github.com/ranjbar-dev/tron-wallet/grpcClient"
	"github.com/ranjbar-dev/tron-wallet/util"
)

const (
	trc20NameSignature           = "0x06fdde03"
	trc20SymbolSignature         = "0x95d89b41"
	trc20DecimalsSignature       = "0x313ce567"
	trc20BalanceOf               = "0x70a08231"
	trc20TransferMethodSignature = "0xa9059cbb"
	trc20FeeLimit                = 10000000
)

type Token struct {
	ContractAddress enums.ContractAddress
}

func (t *Token) GetName(node enums.Node, addressBase58 string) (string, error) {

	c, err := grpcClient.GetGrpcClient(node)
	if err != nil {
		return "", err
	}

	result, err := c.TRC20Call(addressBase58, t.ContractAddress.Base58(), trc20NameSignature, true, 0)
	if err != nil {
		return "", err
	}

	data := util.ToHex(result.GetConstantResult()[0])

	return c.ParseTRC20StringProperty(data)
}

func (t *Token) GetSymbol(node enums.Node, addressBase58 string) (string, error) {

	c, err := grpcClient.GetGrpcClient(node)
	if err != nil {
		return "", err
	}

	result, err := c.TRC20Call(addressBase58, t.ContractAddress.Base58(), trc20SymbolSignature, true, 0)
	if err != nil {
		return "", err
	}

	data := util.ToHex(result.GetConstantResult()[0])

	return c.ParseTRC20StringProperty(data)
}

func (t *Token) GetDecimals(node enums.Node, addressBase58 string) (*big.Int, error) {

	c, err := grpcClient.GetGrpcClient(node)
	if err != nil {
		return nil, err
	}

	result, err := c.TRC20Call(addressBase58, t.ContractAddress.Base58(), trc20DecimalsSignature, true, 0)
	if err != nil {
		return nil, err
	}

	data := util.ToHex(result.GetConstantResult()[0])

	return c.ParseTRC20NumericProperty(data)
}

func (t *Token) GetBalance(node enums.Node, addressBase58 string) (*big.Int, error) {

	address, err := util.Base58ToAddress(addressBase58)
	if err != nil {
		return nil, err
	}

	c, err := grpcClient.GetGrpcClient(node)
	if err != nil {
		return nil, err
	}

	req := trc20BalanceOf + "0000000000000000000000000000000000000000000000000000000000000000"[len(address.Hex())-2:] + address.Hex()[2:]

	result, err := c.TRC20Call(addressBase58, t.ContractAddress.Base58(), req, true, 0)
	if err != nil {
		return nil, err
	}

	data := util.ToHex(result.GetConstantResult()[0])

	r, err := c.ParseTRC20NumericProperty(data)
	if err != nil {
		return nil, fmt.Errorf("contract address %s: %v", t.ContractAddress.Base58(), err)
	}
	if r == nil {
		return nil, fmt.Errorf("contract address %s: invalid balance of %s", t.ContractAddress.Base58(), addressBase58)
	}

	return r, nil
}
