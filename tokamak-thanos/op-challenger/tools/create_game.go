package tools

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/contracts"
	"github.com/tokamak-network/tokamak-thanos/op-service/txmgr"
)

type GameCreator struct {
	contract *contracts.DisputeGameFactoryContract
	txMgr    txmgr.TxManager
}

func NewGameCreator(contract *contracts.DisputeGameFactoryContract, txMgr txmgr.TxManager) *GameCreator {
	return &GameCreator{
		contract: contract,
		txMgr:    txMgr,
	}
}

func (g *GameCreator) CreateGame(ctx context.Context, outputRoot common.Hash, traceType uint64, l2BlockNum uint64) (common.Address, error) {
	txCandidate, err := g.contract.CreateTx(ctx, uint32(traceType), outputRoot, l2BlockNum)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to create tx: %w", err)
	}

	rct, err := g.txMgr.Send(ctx, txCandidate)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to send tx: %w", err)
	}
	if rct.Status != types.ReceiptStatusSuccessful {
		return common.Address{}, fmt.Errorf("game creation transaction (%v) reverted", rct.TxHash.Hex())
	}

	gameAddr, _, _, err := g.contract.DecodeDisputeGameCreatedLog(rct)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to decode game created: %w", err)
	}
	return gameAddr, nil
}
