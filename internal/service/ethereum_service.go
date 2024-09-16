package service

import (
	"context"
	"sync"
	"time"

	"github.com/352174109/trustwallet-homework/internal/dal"
	"github.com/352174109/trustwallet-homework/internal/logs"
	"github.com/352174109/trustwallet-homework/pkg/ethclient"
	"github.com/352174109/trustwallet-homework/pkg/types"
	"github.com/352174109/trustwallet-homework/pkg/utils"
)

type Scanner interface {
	Run()

	Stop() error

	// GetCurrentBlock returns the last scanned block.
	GetCurrentBlock() int
}

type BlockScan struct {
	ctx    context.Context
	cancel context.CancelFunc

	cli            *ethclient.Client
	transactionDal *dal.TransactionDal
	subscribeDal   *dal.SubscribeDal

	parser Parser

	interval         time.Duration
	lastScannedBlock int

	once sync.Once
}

func NewScan(ctx context.Context, transactionDal *dal.TransactionDal, subscribeDal *dal.SubscribeDal, cli *ethclient.Client, startAt int, interval time.Duration) Scanner {
	logs.CtxInfo(ctx, "Blockchain set to start at block: %d", startAt)
	ctx, cancel := context.WithCancel(ctx)
	transactionDal.SetCurrentBlock(ctx, startAt)
	return &BlockScan{
		ctx:    ctx,
		cancel: cancel,

		cli:            cli,
		transactionDal: transactionDal,
		subscribeDal:   subscribeDal,

		interval:         interval,
		lastScannedBlock: startAt,
	}
}

func (b *BlockScan) startScan(ctx context.Context) (int, error) {
	headBlock, err := b.cli.BlockNumber(ctx)
	if err != nil {
		logs.CtxError(ctx, "error querying head block number: %s", err)
		return 0, err
	}

	nextBlockNum := nextBlock(b.lastScannedBlock, headBlock)
	if nextBlockNum == 0 {
		return 0, nil
	}

	block, err := b.scanBlock(ctx, nextBlockNum)
	if err != nil {
		logs.CtxError(ctx, "error scanning block: ", err)
		return 0, err
	}

	err = b.saveBlock(ctx, block.Transactions)
	if err != nil {
		logs.CtxError(ctx, "error saving block: ", err)
	}
	b.lastScannedBlock = nextBlockNum
	b.transactionDal.SetCurrentBlock(ctx, nextBlockNum)

	return b.lastScannedBlock, nil
}

// GetCurrentBlock returns the last scanned block.
func (b *BlockScan) GetCurrentBlock() int {
	return b.lastScannedBlock
}

// Run starts the block scanning process. It will return the number
// of the last scanned block and an error if any. In case of no pending
// blocks to be scanned it will return 0.
func (b *BlockScan) Run() {
	b.once.Do(func() {
		go utils.WrapRecover(b.ctx, b.run)
	})
}

func (b *BlockScan) run() error {
	ticker := time.NewTicker(b.interval)
	for {
		select {
		case <-b.ctx.Done():
			ticker.Stop()
			logs.CtxInfo(b.ctx, "stopping blockchain")
			return nil
		case <-ticker.C:
			ctx := context.Background()
			for scannedBlock, err := b.startScan(ctx); scannedBlock != 0 || err != nil; scannedBlock, err = b.startScan(ctx) {
				if err != nil {
					break
				}
			}
			ticker.Reset(b.interval)
			logs.CtxDebug(ctx, "last scanned block %d\n", b.GetCurrentBlock())
		}
	}
	return nil
}

func (b *BlockScan) Stop() error {
	b.cancel()
	return nil
}

// scanBlock retrieves the block with the given block number and
// returns a map containing the ingoing/outgoing transactions for
// the addresses subscribed.
func (b *BlockScan) scanBlock(ctx context.Context, blockNumber int) (*ethclient.ETHBlock, error) {
	block, err := b.cli.BlockByNumber(ctx, blockNumber)
	if err != nil {
		logs.CtxError(ctx, "error querying block: %s", err.Error())
		return nil, err
	}

	return block, nil
}

func (b *BlockScan) saveBlock(ctx context.Context, transactions []*ethclient.ETHTransaction) error {
	transactionMapByAddr := b.convertToInternalBlock(ctx, transactions)

	for addr, txs := range transactionMapByAddr {
		b.transactionDal.SaveTransaction(ctx, addr, txs)
	}
	return nil
}

// convertToInternalBlock converts a list of ethclient.ETHTransaction into a list of
// types.Transaction.
func (b *BlockScan) convertToInternalBlock(ctx context.Context, txs []*ethclient.ETHTransaction) map[string][]*types.Transaction {
	transactions := make(map[string][]*types.Transaction, len(txs))
	for _, tx := range txs {
		current := &types.Transaction{
			ChainID:     tx.ChainId,
			BlockNumber: tx.BlockNumber,
			Hash:        tx.Hash,
			Nonce:       tx.Nonce,
			From:        tx.From,
			To:          tx.To,
			Value:       tx.Value,
			Gas:         tx.Gas,
			GasPrice:    tx.GasPrice,
			Input:       tx.Input,
		}

		if b.subscribeDal.Subscribed(ctx, current.From) {
			fromTransactions, ok := transactions[current.From]
			if !ok {
				fromTransactions = make([]*types.Transaction, 0)
				transactions[current.From] = fromTransactions
			}
			transactions[current.From] = append(fromTransactions, current)
		}

		if b.subscribeDal.Subscribed(ctx, current.To) {
			toTransactions, ok := transactions[current.To]
			if !ok {
				toTransactions = make([]*types.Transaction, 0)
				transactions[current.To] = toTransactions
			}
			transactions[current.To] = append(toTransactions, current)
		}

	}
	return transactions
}

// nextBlock returns the next block to be scanned. It will return
// 0 if there is any pending block to be scanned. If the last scanned
// block is 0 it will return the head block number.
func nextBlock(lastScannedBlock, headBlock int) int {
	if lastScannedBlock == headBlock {
		return 0
	}
	if lastScannedBlock == 0 {
		return headBlock
	}
	next := lastScannedBlock + 1
	return next
}
