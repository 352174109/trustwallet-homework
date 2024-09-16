package ethclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	ApiVersion           = "2.0"
	GetBlockbusterMethod = "eth_blockNumber"
	GetBlockByNumber     = "eth_getBlockByNumber"
)

type Client struct {
	endpoint string
}

func NewETHClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

// BlockNumber returns the current block number. It will call
// the eth_blockNumber method of the JSON-RPC API in the given endpoint.
func (c *Client) BlockNumber(ctx context.Context) (int, error) {
	body, err := json.Marshal(makeRequestBody(GetBlockbusterMethod, []interface{}{}))
	if err != nil {
		return 0, fmt.Errorf("error marshaling json: %v", err)
	}

	r, err := http.Post(c.endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("error making request: %v", err)
	}
	if r.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error response status code: %v", r.StatusCode)
	}

	getblockNumRsp := &GetBlockNumberResp{}
	if err := json.NewDecoder(r.Body).Decode(getblockNumRsp); err != nil {
		return 0, fmt.Errorf("error decoding response body: %v", err)
	}

	blockNumber, err := strconv.ParseInt(getblockNumRsp.Result[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing response body: %v", err)
	}

	return int(blockNumber), nil
}

func (c *Client) BlockByNumber(ctx context.Context, blockNumber int) (*ETHBlock, error) {
	body, err := json.Marshal(makeRequestBody(GetBlockByNumber, []interface{}{fmt.Sprintf("0x%x", blockNumber), true}))
	if err != nil {
		return nil, fmt.Errorf("error marshaling json: %v", err)
	}

	resp, err := http.Post(c.endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response status code: %v", resp.StatusCode)
	}

	ethBlockRsp := &GetBlockByNumberResp{}
	if err := json.NewDecoder(resp.Body).Decode(ethBlockRsp); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}
	return ethBlockRsp.Result, nil
}

func makeRequestBody(method string, params interface{}) RequestBody {
	return RequestBody{
		Jsonrpc: ApiVersion,
		ID:      1,
		Method:  method,
		Params:  params,
	}
}
