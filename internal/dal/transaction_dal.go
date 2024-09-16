package dal

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/352174109/trustwallet-homework/internal/logs"
	"github.com/352174109/trustwallet-homework/pkg/types"
)

type TransactionDal struct {
	currentBlock int64
	data         map[string][]*types.Transaction

	lock sync.Mutex
}

func NewTransactionDal() (*TransactionDal, error) {
	return &TransactionDal{
		data: make(map[string][]*types.Transaction),
	}, nil
}

func (t *TransactionDal) TransactionByAddr(ctx context.Context, addr string) []*types.Transaction {
	logs.CtxInfo(ctx, "TransactionByAddr addr [%s]", addr)
	t.lock.Lock()
	defer t.lock.Unlock()
	transactions, ok := t.data[addr]
	if !ok {
		return nil
	}

	delete(t.data, addr)
	return transactions
}

func (t *TransactionDal) SaveTransaction(ctx context.Context, addr string, transactions []*types.Transaction) error {
	logs.CtxDebug(ctx, "SaveTransaction addr [%s] transactions number [%d]", addr, len(transactions))
	t.lock.Lock()
	defer t.lock.Unlock()

	olds, ok := t.data[addr]
	if ok {
		olds = append(olds, transactions...)
		t.data[addr] = olds
		return nil
	}

	t.data[addr] = transactions
	return nil
}

func (t *TransactionDal) GetCurrentBlock(ctx context.Context) int {
	return int(atomic.LoadInt64(&t.currentBlock))
}

func (t *TransactionDal) SetCurrentBlock(ctx context.Context, blockNum int) {
	atomic.StoreInt64(&t.currentBlock, int64(blockNum))
}
