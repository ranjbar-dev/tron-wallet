package grpcClient

import (
	"fmt"

	"github.com/ranjbar-dev/tron-wallet/grpcClient/proto/api"
)

func (g *GrpcClient) GetNowBlock() (*api.BlockExtention, error) {
	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetNowBlock2(ctx, new(api.EmptyMessage))

	if err != nil {
		return nil, fmt.Errorf("Get block now: %v", err)
	}

	return result, nil
}

func (g *GrpcClient) GetBlockByNum(num int64) (*api.BlockExtention, error) {
	numMessage := new(api.NumberMessage)
	numMessage.Num = num

	ctx, cancel := g.getContext()
	defer cancel()

	result, err := g.Client.GetBlockByNum2(ctx, numMessage)

	if err != nil {
		return nil, fmt.Errorf("Get block by num: %v", err)

	}
	return result, nil
}
