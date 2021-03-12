module github.com/ovrclk/cosmos-supply-summary

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc2
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.1.1
	github.com/tendermint/tendermint v0.34.8
	google.golang.org/genproto v0.0.0-20210114201628-6edceaf6022f
	google.golang.org/grpc v1.35.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.14.7

replace github.com/cosmos/cosmos-sdk => github.com/ovrclk/cosmos-sdk v0.41.4-akash-2

replace github.com/tendermint/tendermint => github.com/ovrclk/tendermint v0.34.8-akash-1
