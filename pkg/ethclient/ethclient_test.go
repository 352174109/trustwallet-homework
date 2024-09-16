package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req RequestBody
		json.NewDecoder(r.Body).Decode(&req)

		switch req.Method {
		case GetBlockbusterMethod:
			fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":"0x10"}`)
		case GetBlockByNumber:
			fmt.Fprint(w, `{"jsonrpc":"2.0","id":1,"result":{"number":"0x10","hash":"0x1","transactions":[]}}`)
		default:
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	}))
	defer server.Close()

	client := NewETHClient(server.URL)

	t.Run("BlockNumber", func(t *testing.T) {
		blockNumber, err := client.BlockNumber(context.Background())
		if err != nil {
			t.Error(err.Error())
		}
		if blockNumber != 16 {
			t.Error("Block number is not 16")
		}
	})

	t.Run("BlockByNumber", func(t *testing.T) {
		block, err := client.BlockByNumber(context.Background(), 16)
		if err != nil {
			t.Error(err.Error())
		}
		if block.Number != "0x10" {
			t.Error("Block number is not 0x10")
		}

		if block.Hash != "0x1" {
			t.Error("Block hash is not 0x1")
		}
	})
}
