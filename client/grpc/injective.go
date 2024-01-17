package grpc

import (
	"context"

	peggyTypes "bharvest.io/oracle-lens/client/grpc/protobuf/peggy"
	"bharvest.io/oracle-lens/log"
	"google.golang.org/grpc"
)

func NewInjective(host string) *Injective {
	return &Injective{
		host: host,
	}
}

func (client *Injective) Connect(ctx context.Context) error {
	conn, err := grpc.DialContext(
		ctx,
		client.host,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	client.conn = conn
	client.peggyClient = peggyTypes.NewQueryClient(conn)

	log.Info("Injective GRPC connected")

	return nil
}

func (client *Injective) Terminate(_ context.Context) error {
	err := client.conn.Close()
	log.Info("Injective GRPC connection terminated")

	return err
}

func (client *Injective) GetLastPendingValsetSize(ctx context.Context, orch_acc string) (int, error) {
	resp, err := client.peggyClient.LastPendingValsetRequestByAddr(
		ctx,
		&peggyTypes.QueryLastPendingValsetRequestByAddrRequest{Address: orch_acc},
	)
	if err != nil {
		return  0, err
	}

	return resp.Size(), nil
}

func (client *Injective) GetLastPendingBatchSize(ctx context.Context, orch_acc string) (int, error) {
	resp, err := client.peggyClient.LastPendingBatchRequestByAddr(
		ctx,
		&peggyTypes.QueryLastPendingBatchRequestByAddrRequest{Address: orch_acc},
	)
	if err != nil {
		return 0, err
	}

	return resp.Size(), nil
}

func (client *Injective) GetLastEventNonce(ctx context.Context, orch_acc string) (uint64, error) {
	resp, err := client.peggyClient.LastEventByAddr(
		ctx,
		&peggyTypes.QueryLastEventByAddrRequest{Address: orch_acc},
	)
	if err != nil {
		return 0, err
	}

	return resp.GetLastClaimEvent().EthereumEventNonce, nil
}

func (client *Injective) GetLastObservedNonce(ctx context.Context) (uint64, error) {
	resp, err := client.peggyClient.PeggyModuleState(
		ctx,
		&peggyTypes.QueryModuleStateRequest{},
	)
	if err != nil {
		return 0, err
	}

	return resp.State.LastObservedNonce, nil
}

func (client *Injective) GetDelegateKeyByValidator(ctx context.Context, valoper_addr string) (string, error) {
	// valoper_addr => validator operator address
	resp, err := client.peggyClient.GetDelegateKeyByValidator(
		ctx,
		&peggyTypes.QueryDelegateKeysByValidatorAddress{
			ValidatorAddress: valoper_addr,
		},
	)
	if err != nil {
		return "", err
	}

	// resp.OrchestratorAddress => account address
	return resp.OrchestratorAddress, nil
}
