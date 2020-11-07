package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"
	"github.com/spf13/cobra"
)

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
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Summary(context.Background(), &types.QuerySummaryRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res.Supply)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
