package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"
	"github.com/spf13/cobra"
)

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the supply module.
func RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		panic(fmt.Sprintf("couldn't register akash supply grpc routes: %s", err.Error()))
	}
}

// GetQueryCmd returns the query commands for the supply module
func GetQueryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Supply query commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetSummary(),
	)

	return cmd
}

func GetSummary() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Query for supply deployments",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Summary(context.Background(), &types.QuerySummaryRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Supply)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
