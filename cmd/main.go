package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/352174109/trustwallet-homework/internal/dal"
	"github.com/352174109/trustwallet-homework/internal/logs"
	"github.com/352174109/trustwallet-homework/internal/service"
	"github.com/352174109/trustwallet-homework/pkg/ethclient"
)

func main() {
	initialBlock := flag.Int("block", defaultInitialBlock, "block number to start scanning from")
	flag.Parse()

	// Initialize the logger
	logs.SetLevel(context.Background(), logs.LevelInfo)

	subscribeDal, err := dal.NewSubscribeDal()
	if err != nil {
		logs.CtxFatal(context.Background(), err.Error())
		return
	}
	transactionDal, err := dal.NewTransactionDal()
	if err != nil {
		logs.CtxFatal(context.Background(), err.Error())
		return
	}
	parser, err := service.NewEthereumParser(subscribeDal, transactionDal)
	if err != nil {
		logs.CtxFatal(context.Background(), err.Error())
		return
	}
	
	srv, err := service.NewService(context.Background(), parser)
	if err != nil {
		logs.CtxFatal(context.Background(), err.Error())
		return
	}
	// Start the service, receive command line arguments
	srv.Start(context.Background())

	ethCli := ethclient.NewETHClient(endPoint)
	scanService := service.NewScan(context.Background(), transactionDal, subscribeDal, ethCli, *initialBlock, time.Second*10)
	// Start blockchain service, pull block transactions information every 10 seconds
	scanService.Run()

	// Mock data
	mockData(transactionDal, subscribeDal)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalChannel
	logs.CtxInfo(context.Background(), "\nReceived signal: %s. Shutting down...\n", sig)

	srv.Stop()
	scanService.Stop()
}
