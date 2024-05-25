package grpcClient

import (
	"fmt"

	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/api"
	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/core"
	"github.com/ranjbar-dev/tron-wallet/util"
	"google.golang.org/protobuf/proto"
)

func (g *GrpcClient) Transfer(from, toAddress string, amount int64) (*api.TransactionExtention, error) {
	var err error

	contract := &core.TransferContract{}
	if contract.OwnerAddress, err = util.DecodeCheck(from); err != nil {
		return nil, err
	}
	if contract.ToAddress, err = util.DecodeCheck(toAddress); err != nil {
		return nil, err
	}
	contract.Amount = amount

	ctx, cancel := g.getContext()
	defer cancel()

	tx, err := g.Client.CreateTransaction2(ctx, contract)
	if err != nil {
		return nil, err
	}
	if proto.Size(tx) == 0 {
		return nil, fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return nil, fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return tx, nil
}

func (g *GrpcClient) Broadcast(tx *core.Transaction) (*api.Return, error) {
	ctx, cancel := g.getContext()
	defer cancel()
	result, err := g.Client.BroadcastTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	if !result.GetResult() {
		return result, fmt.Errorf("result error: %s", result.GetMessage())
	}
	if result.GetCode() != api.Return_SUCCESS {
		return result, fmt.Errorf("result error(%s): %s", result.GetCode(), result.GetMessage())
	}
	return result, nil
}
