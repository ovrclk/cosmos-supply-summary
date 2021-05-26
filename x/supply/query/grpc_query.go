package query

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"
)

// Querier creates supply querier instance with account, banking, distribution and staking keeperss
type Querier struct {
	accKeeper   types.AccountKeeper
	bankKeeper  types.BankKeeper
	distrKeeper types.DistrKeeper
	stKeeper    types.StakingKeeper
}

// NewQuerier return new supply querier instance
func NewQuerier(accKeeper types.AccountKeeper, bankKeeper types.BankKeeper, stKeeper types.StakingKeeper,
	distrKeeper types.DistrKeeper) Querier {
	return Querier{
		accKeeper:   accKeeper,
		bankKeeper:  bankKeeper,
		stKeeper:    stKeeper,
		distrKeeper: distrKeeper,
	}
}

var _ types.QueryServer = Querier{}

// Summary returns summary of supply
func (q Querier) Summary(c context.Context, req *types.QuerySummaryRequest) (*types.QuerySummaryResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var supplyData types.Supply
	supplyData.Total = q.bankKeeper.GetSupply(ctx).GetTotal()
	bondDenom := q.stKeeper.BondDenom(ctx)

	delegationsMap := make(map[string]sdk.Coins)
	q.stKeeper.IterateAllDelegations(ctx, func(delegation stakingtypes.Delegation) bool {
		// Converting delegated shares to sdk.Coin
		delegated := sdk.NewCoin(bondDenom, delegation.Shares.TruncateInt())
		delegationsMap[delegation.DelegatorAddress] = delegationsMap[delegation.DelegatorAddress].Add(delegated)
		return false
	})
	q.accKeeper.IterateAccounts(ctx, func(account authtypes.AccountI) bool {
		if ma, ok := account.(*authtypes.ModuleAccount); ok {
			switch ma.Name {
			case stakingtypes.NotBondedPoolName, stakingtypes.BondedPoolName:
				return false
			}
		}
		delegatedTokens := delegationsMap[account.GetAddress().String()]
		balances := q.bankKeeper.GetAllBalances(ctx, account.GetAddress())
		va, ok := account.(vestingexported.VestingAccount)
		if !ok {
			supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedTokens...)
			supplyData.Available.Unbonded = supplyData.Available.Unbonded.Add(balances...)
		} else {
			vestingCoins := va.GetVestingCoins(ctx.BlockTime())
			delegatedVesting := va.GetDelegatedVesting()
			lockedCoins := va.LockedCoins(ctx.BlockTime())
			var spendableCoins sdk.Coins
			if balances.AmountOf(bondDenom).GT(lockedCoins.AmountOf(bondDenom)) {
				spendableCoins = balances.Sub(lockedCoins)
			}
			if delegatedVesting.AmountOf(bondDenom).GT(vestingCoins.AmountOf(bondDenom)) {
				supplyData.Vesting.Bonded = supplyData.Vesting.Bonded.Add(vestingCoins...)
				supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedVesting...).Sub(vestingCoins)
			} else {
				supplyData.Vesting.Bonded = supplyData.Vesting.Bonded.Add(delegatedVesting...)
				supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedTokens...).Sub(delegatedVesting)
			}
			supplyData.Vesting.Unbonded = supplyData.Vesting.Unbonded.Add(lockedCoins...)
			supplyData.Available.Unbonded = supplyData.Available.Unbonded.Add(spendableCoins...)
		}
		return false
	})

	communityPool, _ := q.distrKeeper.GetFeePoolCommunityCoins(ctx).TruncateDecimal()

	supplyData.Circulating = supplyData.Total.Sub(supplyData.Vesting.Unbonded).Sub(supplyData.Vesting.Bonded).Sub(communityPool)

	return &types.QuerySummaryResponse{
		Supply: supplyData,
		Height: ctx.BlockHeight(),
	}, nil
}
