package rest

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/ovrclk/cosmos-supply-summary/x/supply/query"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	decimals float64 = 6
	denom            = "uakt"
)

// RegisterRoutes registers all query routes
func RegisterRoutes(ctx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/supply/summary", circulatingSupplyHandler(ctx)).Methods("GET")
}

func circulatingSupplyHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if len(q) == 0 {
			res, err := query.NewRawClient(ctx, types.ModuleName).CirculatingSupply()
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			rest.PostProcessResponse(w, ctx, res)
		} else {
			res, err := query.NewClient(ctx, types.ModuleName).CirculatingSupply()
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			metric := sdk.Coins{}
			if q == "total" {
				metric = res.Total
			} else if q == "circulating" {
				metric = res.Circulating
			} else {
				rest.WriteErrorResponse(w, http.StatusBadRequest, "wrong q value")
				return
			}
			supply := sdk.Coin{}
			for _, coin := range metric {
				if coin.Denom == denom {
					supply = coin
				}
			}
			if supply.Denom == "" {
				rest.WriteErrorResponse(w, http.StatusBadRequest, "no supply found with default denom")
				return
			}
			fmt.Fprintf(w, "%.6f", float64(supply.Amount.Int64())/math.Pow(10, decimals))
		}
	}
}
