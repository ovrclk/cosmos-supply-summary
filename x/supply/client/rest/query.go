package rest

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterRoutes registers all query routes
func RegisterRoutes(ctx client.Context, r *mux.Router) {
	r.HandleFunc("/supply/summary", circulatingSupplyHandler(ctx)).Methods("GET")
}

func circulatingSupplyHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		queryClient := types.NewQueryClient(ctx)
		res, err := queryClient.Summary(context.Background(), &types.QuerySummaryRequest{})
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(q) == 0 {
			out, err := ctx.JSONMarshaler.MarshalJSON(res)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			rest.PostProcessResponse(w, ctx, out)
			return
		}

		metric := sdk.Coins{}
		switch q {
		case "total":
			metric = res.Supply.Total
		case "circulating":
			metric = res.Supply.Circulating
		default:
			rest.WriteErrorResponse(w, http.StatusBadRequest, "wrong query value, allowed values are `total`, `circulating`")
			return
		}

		denom := r.URL.Query().Get("denom")
		decString := r.URL.Query().Get("decimals")
		if len(denom) == 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "`denom` value is required")
			return
		}
		if len(decString) == 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "`decimals` value is required")
			return
		}

		decimals, err := strconv.ParseUint(decString, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		supply := sdk.Coin{}
		for _, coin := range metric {
			if coin.Denom == denom {
				supply = coin
			}
		}
		if supply.Denom == "" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "supply not found with given denom")
			return
		}
		fmt.Fprintf(w, "%.6f", float64(supply.Amount.Int64())/math.Pow(10, float64(decimals)))
	}
}
