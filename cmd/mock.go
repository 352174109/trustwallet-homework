package main

import (
	"context"

	"github.com/352174109/trustwallet-homework/internal/dal"
	"github.com/352174109/trustwallet-homework/pkg/types"
)

func mockData(transactionDal *dal.TransactionDal, subscribeDal *dal.SubscribeDal) {
	ctx := context.Background()
	subscribeDal.Subscribe(ctx, "0xe75ed6f453c602bd696ce27af11565edc9b46b0d")
	subscribeDal.Subscribe(ctx, "0x00000000009e50a7ddb7a7b0e2ee6604fd120e49")

	transactions := []*types.Transaction{
		{
			ChainID:     "0x1",
			BlockNumber: "0x13cd296",
			Hash:        "0x1b0d6db9bee0b358beda0da81de82cec9ebba5b4488f460943b979fe2315f3c3",
			Nonce:       "0x2c08b",
			From:        "0xe75ed6f453c602bd696ce27af11565edc9b46b0d",
			To:          "0x00000000009e50a7ddb7a7b0e2ee6604fd120e49",
			Value:       "0xf5232269",
			Gas:         "0x46b47",
			GasPrice:    "0x1a5c9d8f9",
			Input:       "0x960d1f9afe7a4e6c6aa2f928b71a512b2e6644d7a7e5593d148b89b41a0889322bba387c825180ebfb62bd8e6969ebe5b5e52d02aa1efb3c159d81db1c006d",
		},
	}

	transactionDal.SaveTransaction(ctx, "0xe75ed6f453c602bd696ce27af11565edc9b46b0d", transactions)
	transactionDal.SaveTransaction(ctx, "0x00000000009e50a7ddb7a7b0e2ee6604fd120e49", transactions)
}
