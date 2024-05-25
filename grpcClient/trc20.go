package grpcClient

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"unicode/utf8"

	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/api"
	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/core"
	"github.com/ranjbar-dev/tron-wallet/util"
	"google.golang.org/protobuf/proto"
)

// TRC20Call make constant call
func (g *GrpcClient) TRC20Call(fromAddressBase58 string, contractAddressBase58 string, data string, constant bool, feeLimit int64) (*api.TransactionExtention, error) {

	var err error
	fromAddress := util.HexToAddress("410000000000000000000000000000000000000000")
	if len(fromAddressBase58) > 0 {
		fromAddress, err = util.Base58ToAddress(fromAddressBase58)
		if err != nil {
			return nil, err
		}
	}

	dataBytes, err := util.FromHex(data)
	if err != nil {
		return nil, err
	}

	contractAddress, err := util.Base58ToAddress(contractAddressBase58)
	if err != nil {
		return nil, err
	}

	ct := &core.TriggerSmartContract{
		OwnerAddress:    fromAddress.Bytes(),
		ContractAddress: contractAddress.Bytes(),
		Data:            dataBytes,
	}
	result := &api.TransactionExtention{}
	if constant {
		result, err = g.triggerConstantContract(ct)

	} else {
		result, err = g.triggerContract(ct, feeLimit)
	}
	if err != nil {
		return nil, err
	}
	if result.Result.Code > 0 {
		return result, fmt.Errorf(string(result.Result.Message))
	}
	return result, nil

}

// triggerConstantContract and return tx result
func (g *GrpcClient) triggerConstantContract(ct *core.TriggerSmartContract) (*api.TransactionExtention, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	return g.Client.TriggerConstantContract(ctx, ct)
}

// triggerContract and return tx result
func (g *GrpcClient) triggerContract(ct *core.TriggerSmartContract, feeLimit int64) (*api.TransactionExtention, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.TriggerContract(ctx, ct)
	if err != nil {
		return nil, err
	}

	if tx.Result.Code > 0 {
		return nil, fmt.Errorf("%s", string(tx.Result.Message))
	}
	if feeLimit > 0 {
		tx.Transaction.RawData.FeeLimit = feeLimit
		// update hash
		g.UpdateHash(tx)
	}
	return tx, err
}

// UpdateHash after local changes
func (g *GrpcClient) UpdateHash(tx *api.TransactionExtention) error {
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		return err
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	tx.Txid = hash
	return nil
}

// ParseTRC20StringProperty get string from data
func (g *GrpcClient) ParseTRC20StringProperty(data string) (string, error) {
	if util.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) > 128 {
		n, _ := g.ParseTRC20NumericProperty(data[64:128])
		if n != nil {
			l := n.Uint64()
			if 2*int(l) <= len(data)-128 {
				b, err := hex.DecodeString(data[128 : 128+2*l])
				if err == nil {
					return string(b), nil
				}
			}
		}
	} else if len(data) == 64 {
		// allow string properties as 32 bytes of UTF-8 data
		b, err := hex.DecodeString(data)
		if err == nil {
			i := bytes.Index(b, []byte{0})
			if i > 0 {
				b = b[:i]
			}
			if utf8.Valid(b) {
				return string(b), nil
			}
		}
	}
	return "", fmt.Errorf("Cannot parse %s,", data)
}

// ParseTRC20NumericProperty get number from data
func (g *GrpcClient) ParseTRC20NumericProperty(data string) (*big.Int, error) {
	if util.Has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) == 64 {
		var n big.Int
		_, ok := n.SetString(data, 16)
		if ok {
			return &n, nil
		}
	}
	return nil, fmt.Errorf("Cannot parse %s", data)
}
