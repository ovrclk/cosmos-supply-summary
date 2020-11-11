package rest

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/ovrclk/cosmos-supply-summary/x/supply/query"
	"github.com/ovrclk/cosmos-supply-summary/x/supply/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
			return
		}
		denom := r.URL.Query().Get("denom")
		decString := r.URL.Query().Get("decimal")
		if len(denom) == 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "`denom` value is required")
			return
		}
		if len(decString) == 0 {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "`decimal` value is required")
			return
		}

		decimals, err := strconv.ParseFloat(decString, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

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
			rest.WriteErrorResponse(w, http.StatusBadRequest, "wrong query value, allowed values are `total`, `circulating`")
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
		fmt.Fprintf(w, "%.6f", float64(supply.Amount.Int64())/math.Pow(10, decimals))
	}
}
