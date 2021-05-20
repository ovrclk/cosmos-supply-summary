package rest

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
)

// CurrentBlockHeight returns current block height of node
func CurrentBlockHeight(ctx client.Context) (int64, error) {
	client, err := ctx.GetNode()
	if err != nil {
		return 0, err
	}
	status, err := client.Status(context.Background())
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}
