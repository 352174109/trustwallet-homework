package service

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/352174109/trustwallet-homework/internal/logs"
	"github.com/352174109/trustwallet-homework/pkg/utils"
)

type Service struct {
	ctx     context.Context
	cancel  context.CancelFunc
	running int32

	parser Parser

	once sync.Once
}

func NewService(ctx context.Context, parser Parser) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	return &Service{
		ctx:    ctx,
		cancel: cancel,

		parser: parser,
	}, nil
}

func (s *Service) Start(ctx context.Context) {
	s.once.Do(func() {
		atomic.StoreInt32(&s.running, 1)

		logs.CtxInfo(ctx, "Service started.")

		go utils.WrapRecover(ctx, s.listenForCommands)

		logs.CtxInfo(ctx, "Service started successful.")
	})
}

func (s *Service) Stop() {
	atomic.StoreInt32(&s.running, 0)
	s.cancel()
}

// monitory user commands
func (s *Service) listenForCommands() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-s.ctx.Done():
			logs.CtxInfo(s.ctx, "stopping console service")
			return nil
		default:
			fmt.Print("> ")
			command, _ := reader.ReadString('\n')
			command = strings.TrimSpace(command)

			s.handleCommand(command)
		}
	}

	return nil
}

func (s *Service) handleCommand(input string) {
	currentCtx := context.Background()
	logs.CtxInfo(currentCtx, "Received command: %s", input)

	args := strings.Fields(input)

	if len(args) < 1 {
		fmt.Println("No command provided.")
		return
	}

	switch args[0] {
	case "getCurrentBlock":
		block := s.parser.GetCurrentBlock(currentCtx)
		logs.CtxInfo(currentCtx, "Current Block: %d", block)
	case "subscribe":
		if len(args) != 2 {
			logs.CtxInfo(currentCtx, "Invalid number of arguments. Usage: subscribe <address>")
			return
		}
		success := s.parser.Subscribe(currentCtx, args[1])
		if success {
			logs.CtxInfo(currentCtx, "Subscribed to address: %s", args[1])
		} else {
			logs.CtxInfo(currentCtx, "Failed to subscribe to address: %s", args[1])
		}
	case "getTransactions":
		if len(args) != 2 {
			logs.CtxInfo(currentCtx, "Invalid number of arguments. Usage: getTransactions <address>")
			return
		}
		transactions := s.parser.GetTransactions(currentCtx, args[1])
		if len(transactions) > 0 {
			logs.CtxInfo(currentCtx, "Transactions:")
			for _, transaction := range transactions {
				transactionStr, _ := json.Marshal(transaction)
				logs.CtxInfo(currentCtx, "- %+s", transactionStr)
			}
		} else {
			logs.CtxInfo(currentCtx, "No transactions found for address: %s", args[1])
		}
	case "help":
		printHelp()
	default:
		logs.CtxInfo(currentCtx, "Unknown command: '%s'. Use 'help' to see available commands.\n", input)
		printHelp()
	}
}

// 打印帮助信息
func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  getCurrentBlock            - Subscribed the latest block number")
	fmt.Println("  subscribe <address>        - Subscribe to monitor a specific address")
	fmt.Println("  getTransactions <address>  - Subscribed transactions related to a specific address")
	fmt.Println("  help                       - Show available commands and usage")
}
