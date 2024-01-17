package grpc

import (
	"context"

	peggyTypes "bharvest.io/oracle-lens/client/grpc/protobuf/peggy"
	umeeOracleTypes "bharvest.io/oracle-lens/client/grpc/protobuf/umee-oracle"
	"google.golang.org/grpc"
)

type Client interface {
	Connect(ctx context.Context) error
	Terminate(_ context.Context) error
}

type (
	Injective struct {
		host string
		conn *grpc.ClientConn
		peggyClient peggyTypes.QueryClient
	}
	Umee struct {
		host string
		conn *grpc.ClientConn
		oracleClient umeeOracleTypes.QueryClient
	}
)

type (
	UmeeOracleParams struct {
		AcceptList map[string]bool
		SlashWindow uint64
		MinUptime float64
	}
)
