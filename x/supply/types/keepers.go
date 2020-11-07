package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankexported "github.com/cosmos/cosmos-sdk/x/bank/exported"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper for module accounts
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	IterateAccounts(ctx sdk.Context, process func(authtypes.AccountI) bool)
}

// BankKeeper defines the expected bank kseeper for module accounts
type BankKeeper interface {
	GetSupply(ctx sdk.Context) (supply bankexported.SupplyI)
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// StakingKeeper defines the expected staking keeper for module accounts
type StakingKeeper interface {
	IterateAllDelegations(ctx sdk.Context, cb func(stakingtypes.Delegation) bool)
	BondDenom(ctx sdk.Context) string
}

// DistrKeeper defines the expected distribution keeper for module accounts
type DistrKeeper interface {
	GetFeePoolCommunityCoins(ctx sdk.Context) sdk.DecCoins
}
