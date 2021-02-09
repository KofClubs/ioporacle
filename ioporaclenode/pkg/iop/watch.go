package iop

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	log "github.com/sirupsen/logrus"
)

func (n *OracleNode) watch(ctx context.Context) {
	go func() {
		if err := n.dkg.Listen(ctx); err != nil {
			log.Errorf("watch register oracle node log: %v", err)
		}
	}()
	go func() {
		if err := n.watchRegisterOracleNodeLog(ctx); err != nil {
			log.Errorf("watch register oracle node log: %v", err)
		}
	}()
	go func() {
		if err := n.watchValidateTransactionLog(ctx); err != nil {
			log.Errorf("watch verify transaction log: %v", err)
		}
	}()
}

func (n *OracleNode) watchRegisterOracleNodeLog(ctx context.Context) error {
	sink := make(chan *RegistryContractRegisterOracleNodeLog)
	defer close(sink)

	sub, err := n.registryContract.WatchRegisterOracleNodeLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received register oracle node event %s", event.Sender.String())
			if err = n.handleRegisterOracleNodeLog(event); err != nil {
				log.Errorf("handle register oracle node log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (n *OracleNode) watchValidateTransactionLog(ctx context.Context) error {
	sink := make(chan *OracleContractValidateTransactionLog)
	defer close(sink)

	sub, err := n.oracleContract.WatchValidateTransactionLog(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			log.Infof("Received verify transaction event %s", event.Id)
			isAggregator, err := n.registryContract.IsAggregator(nil, n.account)
			if err != nil {
				log.Errorf("is aggregator: %v", err)
				continue
			}
			if !isAggregator {
				continue
			}
			if err := n.handleVerifyTransactionLog(ctx, event); err != nil {
				log.Errorf("handle verify transaction log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
