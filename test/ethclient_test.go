package test

import (
	"context"
	"testing"

	"github.com/352174109/trustwallet-homework/pkg/ethclient"
)

const (
	endPoint = "https://cloudflare-eth.com"
)

func Test_BlockNumber(t *testing.T) {
	cli := ethclient.NewETHClient(endPoint)
	number, err := cli.BlockNumber(context.Background())
	t.Log(err)
	t.Log(number)
}
