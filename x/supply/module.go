package supply

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/ovrclk/cosmos-supply-summary/x/supply/client/cli"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/query"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the supply module.
type AppModuleBasic struct {
	cdc codec.Marshaler
}

// Name returns supply module's name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the module types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

// RegisterInterfaces registers the module's interface types
func (AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

// DefaultGenesis returns default genesis state as raw bytes for the supply module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONMarshaler) json.RawMessage {
	return nil
}

// ValidateGenesis validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(_ codec.JSONMarshaler, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

// RegisterRESTRoutes registers rest routes for this module
func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the supply module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		panic(fmt.Sprintf("couldn't register deployment grpc routes: %s", err.Error()))
	}
}

// GetQueryCmd returns the root query command of this module
// This module has a query which is directly added to SDK supply queries
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// GetTxCmd returns the root tx command of this module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil
}

// AppModule implements an application module for the supply module.
type AppModule struct {
	AppModuleBasic

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	StakingKeeper types.StakingKeeper
	DistrKeeper   types.DistrKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Marshaler, accKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
	stKeeper types.StakingKeeper, distrKeeper types.DistrKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		AccountKeeper:  accKeeper,
		BankKeeper:     bankKeeper,
		StakingKeeper:  stKeeper,
		DistrKeeper:    distrKeeper,
	}
}

// Name returns the supply module name
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers module invariant
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the supply module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.ModuleName, nil)
}

// QuerierRoute returns the supply module's querier route name.
func (am AppModule) QuerierRoute() string {
	return types.ModuleName
}

// LegacyQuerierHandler returns the sdk.Querier for supply module
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers the module's services
func (am AppModule) RegisterServices(cfg module.Configurator) {
	querier := query.NewQuerier(am.AccountKeeper, am.BankKeeper, am.StakingKeeper, am.DistrKeeper)
	types.RegisterQueryServer(cfg.QueryServer(), querier)
}

// BeginBlock performs no-op
func (am AppModule) BeginBlock(ctx sdk.Context, beginBlock abci.RequestBeginBlock) {}

// EndBlock performs no-op
func (am AppModule) EndBlock(ctx sdk.Context, endBlock abci.RequestEndBlock) []abci.ValidatorUpdate {
	return nil
}

// InitGenesis performs genesis initialization for the supply module. It returns
// no validator updates.
func (am AppModule) InitGenesis(_ sdk.Context, _ codec.JSONMarshaler, _ json.RawMessage) []abci.ValidatorUpdate {
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the supply module.
func (am AppModule) ExportGenesis(_ sdk.Context, _ codec.JSONMarshaler) json.RawMessage {
	return nil
}
