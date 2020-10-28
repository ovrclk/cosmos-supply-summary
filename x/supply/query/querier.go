package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	vestingexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"

	"github.com/ovrclk/cosmos-supply-summary/sdkutil"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates and returns a new supply querier instance
func NewQuerier(cdc *codec.Codec, accKeeper types.AccountKeeper, supKeeper types.SupplyKeeper,
	stKeeper types.StakingKeeper, distrKeeper types.DistrKeeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case circulatingPath:
			return queryCirculatingSupply(ctx, cdc, accKeeper, supKeeper, stKeeper, distrKeeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown query for endpoint: %s", path[0])
		}
	}
}

func queryCirculatingSupply(ctx sdk.Context, cdc *codec.Codec, accKeeper types.AccountKeeper, supKeeper types.SupplyKeeper,
	stKeeper types.StakingKeeper, distrKeeper types.DistrKeeper) (res []byte, err error) {
	var supplyData Supply
	supplyData.Total = supKeeper.GetSupply(ctx).GetTotal()
	bondDenom := stKeeper.BondDenom(ctx)

	delegationsMap := make(map[string]sdk.Coins)
	stKeeper.IterateAllDelegations(ctx, func(delegation staking.Delegation) bool {
		// Converting delegated shares to sdk.Coin
		delegated := sdk.NewCoin(bondDenom, delegation.Shares.TruncateInt())
		delegationsMap[delegation.DelegatorAddress.String()] = delegationsMap[delegation.DelegatorAddress.String()].Add(delegated)
		return false
	})
	accKeeper.IterateAccounts(ctx, func(account exported.Account) bool {
		if ma, ok := account.(*supply.ModuleAccount); ok {
			switch ma.Name {
			case staking.NotBondedPoolName, staking.BondedPoolName:
				return false
			}
		}
		delegatedTokens := delegationsMap[account.GetAddress().String()]
		va, ok := account.(vestingexported.VestingAccount)
		if !ok {
			supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedTokens...)
			supplyData.Available.Unbonded = supplyData.Available.Unbonded.Add(account.GetCoins()...)
		} else {
			vestingCoins := va.GetVestingCoins(ctx.BlockTime())
			delegatedVesting := va.GetDelegatedVesting()
			if delegatedVesting.AmountOf(bondDenom).GT(vestingCoins.AmountOf(bondDenom)) {
				supplyData.Vesting.Bonded = supplyData.Vesting.Bonded.Add(vestingCoins...)
				supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedVesting...).Sub(vestingCoins)
			} else {
				supplyData.Vesting.Bonded = supplyData.Vesting.Bonded.Add(delegatedVesting...)
				supplyData.Available.Bonded = supplyData.Available.Bonded.Add(delegatedTokens...).Sub(delegatedVesting)
			}
			supplyData.Vesting.Unbonded = supplyData.Vesting.Unbonded.Add(va.GetCoins()...).Sub(va.SpendableCoins(ctx.BlockTime()))
			supplyData.Available.Unbonded = supplyData.Available.Unbonded.Add(va.SpendableCoins(ctx.BlockTime())...)
		}
		return false
	})

	communityPool, _ := distrKeeper.GetFeePoolCommunityCoins(ctx).TruncateDecimal()

	supplyData.Circulating = supplyData.Total.Sub(supplyData.Vesting.Unbonded).Sub(supplyData.Vesting.Bonded).Sub(communityPool)
	return sdkutil.RenderQueryResponse(cdc, supplyData)
}
