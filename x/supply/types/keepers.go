package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
)

// AccountKeeper defines the expected account keeper for module accounts
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account
	IterateAccounts(ctx sdk.Context, process func(exported.Account) bool)
}

// SupplyKeeper defines the expected supply keeper for module accounts
type SupplyKeeper interface {
	GetSupply(ctx sdk.Context) (supply supplyexported.SupplyI)
}

// StakingKeeper defines the expected staking keeper for module accounts
type StakingKeeper interface {
	IterateAllDelegations(ctx sdk.Context, cb func(stakingtypes.Delegation) bool)
	BondDenom(ctx sdk.Context) string
}
