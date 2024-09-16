package service

import (
	"context"

	"github.com/352174109/trustwallet-homework/internal/dal"
	"github.com/352174109/trustwallet-homework/internal/logs"
	"github.com/352174109/trustwallet-homework/pkg/types"
)

type Parser interface {
	// GetCurrentBlock last parsed block
	GetCurrentBlock(ctx context.Context) int
	// Subscribe add address to observer
	Subscribe(ctx context.Context, address string) bool
	// GetTransactions list of inbound or outbound transactions for an address
	GetTransactions(ctx context.Context, address string) []*types.Transaction
}

// EthereumParser implements the Parser interface
type EthereumParser struct {
	subscribeDal   *dal.SubscribeDal
	transactionDal *dal.TransactionDal
}

// NewEthereumParser creates a new EthereumParser instance
func NewEthereumParser(subscribeDal *dal.SubscribeDal, transactionDal *dal.TransactionDal) (Parser, error) {
	return &EthereumParser{
		subscribeDal:   subscribeDal,
		transactionDal: transactionDal,
	}, nil
}

// GetCurrentBlock returns the last parsed block number
func (p *EthereumParser) GetCurrentBlock(ctx context.Context) int {
	return p.transactionDal.GetCurrentBlock(ctx)
}

// Subscribe adds an address to the list of subscribed addresses for monitoring
func (p *EthereumParser) Subscribe(ctx context.Context, address string) bool {
	if err := p.subscribeDal.Subscribe(ctx, address); err != nil {
		logs.CtxError(ctx, "Subscribed to address: %s, err: %s", address, err.Error())
		return false
	}
	logs.CtxInfo(ctx, "Address [%s] subscribed successful", address)
	return true
}

// GetTransactions returns the list of transactions (inbound/outbound) for a given address
// if not subscribed, then subscribe this address and return nil
func (p *EthereumParser) GetTransactions(ctx context.Context, address string) []*types.Transaction {
	if address == "" {
		logs.CtxWarn(ctx, "Address already subscribed: %s", address)
		return nil
	}
	if !p.subscribeDal.Subscribed(ctx, address) {
		err := p.subscribeDal.Subscribe(ctx, address)
		if err != nil {
			logs.CtxWarn(ctx, err.Error())
		}
		return nil
	}

	return p.transactionDal.TransactionByAddr(ctx, address)
}
