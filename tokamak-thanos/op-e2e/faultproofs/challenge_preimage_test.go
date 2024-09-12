package faultproofs

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	op_e2e "github.com/tokamak-network/tokamak-thanos/op-e2e"
	"github.com/tokamak-network/tokamak-thanos/op-e2e/e2eutils/challenger"
	"github.com/tokamak-network/tokamak-thanos/op-e2e/e2eutils/disputegame"
	"github.com/tokamak-network/tokamak-thanos/op-e2e/e2eutils/disputegame/preimage"
)

func TestChallengeLargePreimages_ChallengeFirst(t *testing.T) {
	op_e2e.InitParallel(t, op_e2e.UsesCannon)
	ctx := context.Background()
	sys, _ := StartFaultDisputeSystem(t)
	t.Cleanup(sys.Close)

	disputeGameFactory := disputegame.NewFactoryHelper(t, ctx, sys)
	disputeGameFactory.StartChallenger(ctx, "Challenger",
		challenger.WithAlphabet(),
		challenger.WithPrivKey(sys.Cfg.Secrets.Alice))
	preimageHelper := disputeGameFactory.PreimageHelper(ctx)
	ident := preimageHelper.UploadLargePreimage(ctx, preimage.MinPreimageSize,
		preimage.WithReplacedCommitment(0, common.Hash{0xaa}))

	require.NotEqual(t, ident.Claimant, common.Address{})

	preimageHelper.WaitForChallenged(ctx, ident)
}

func TestChallengeLargePreimages_ChallengeMiddle(t *testing.T) {
	op_e2e.InitParallel(t, op_e2e.UsesCannon)
	ctx := context.Background()
	sys, _ := StartFaultDisputeSystem(t)
	t.Cleanup(sys.Close)
	disputeGameFactory := disputegame.NewFactoryHelper(t, ctx, sys)
	disputeGameFactory.StartChallenger(ctx, "Challenger",
		challenger.WithAlphabet(),
		challenger.WithPrivKey(sys.Cfg.Secrets.Mallory))
	preimageHelper := disputeGameFactory.PreimageHelper(ctx)
	ident := preimageHelper.UploadLargePreimage(ctx, preimage.MinPreimageSize,
		preimage.WithReplacedCommitment(10, common.Hash{0xaa}))

	require.NotEqual(t, ident.Claimant, common.Address{})

	preimageHelper.WaitForChallenged(ctx, ident)
}

func TestChallengeLargePreimages_ChallengeLast(t *testing.T) {
	op_e2e.InitParallel(t, op_e2e.UsesCannon)
	ctx := context.Background()
	sys, _ := StartFaultDisputeSystem(t)
	t.Cleanup(sys.Close)
	disputeGameFactory := disputegame.NewFactoryHelper(t, ctx, sys)
	disputeGameFactory.StartChallenger(ctx, "Challenger",
		challenger.WithAlphabet(),
		challenger.WithPrivKey(sys.Cfg.Secrets.Mallory))
	preimageHelper := disputeGameFactory.PreimageHelper(ctx)
	ident := preimageHelper.UploadLargePreimage(ctx, preimage.MinPreimageSize,
		preimage.WithLastCommitment(common.Hash{0xaa}))

	require.NotEqual(t, ident.Claimant, common.Address{})

	preimageHelper.WaitForChallenged(ctx, ident)
}
