package ethclient

type GetBlockByNumberResp struct {
	Jsonrpc string    `json:"jsonrpc"`
	Result  *ETHBlock `json:"result"`
	Id      int       `json:"id"`
}

type ETHBlock struct {
	BaseFeePerGas         string            `json:"baseFeePerGas"`
	BlobGasUsed           string            `json:"blobGasUsed"`
	Difficulty            string            `json:"difficulty"`
	ExcessBlobGas         string            `json:"excessBlobGas"`
	ExtraData             string            `json:"extraData"`
	GasLimit              string            `json:"gasLimit"`
	GasUsed               string            `json:"gasUsed"`
	Hash                  string            `json:"hash"`
	LogsBloom             string            `json:"logsBloom"`
	Miner                 string            `json:"miner"`
	MixHash               string            `json:"mixHash"`
	Nonce                 string            `json:"nonce"`
	Number                string            `json:"number"`
	ParentBeaconBlockRoot string            `json:"parentBeaconBlockRoot"`
	ParentHash            string            `json:"parentHash"`
	ReceiptsRoot          string            `json:"receiptsRoot"`
	Sha3Uncles            string            `json:"sha3Uncles"`
	Size                  string            `json:"size"`
	StateRoot             string            `json:"stateRoot"`
	Timestamp             string            `json:"timestamp"`
	TotalDifficulty       string            `json:"totalDifficulty"`
	Transactions          []*ETHTransaction `json:"transactions"`
	TransactionsRoot      string            `json:"transactionsRoot"`
	Uncles                []interface{}     `json:"uncles"`
	Withdrawals           []*ETHWithdrawal  `json:"result"`
	WithdrawalsRoot       string            `json:"withdrawalsRoot"`
}

type ETHTransaction struct {
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	From                 string `json:"from"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	Hash                 string `json:"hash"`
	Input                string `json:"input"`
	Nonce                string `json:"nonce"`
	To                   string `json:"to"`
	TransactionIndex     string `json:"transactionIndex"`
	Value                string `json:"value"`
	Type                 string `json:"type"`
	AccessList           []struct {
		Address     string   `json:"address"`
		StorageKeys []string `json:"storageKeys"`
	} `json:"accessList"`
	ChainId string `json:"chainId"`
	V       string `json:"v"`
	R       string `json:"r"`
	S       string `json:"s"`
}

type ETHWithdrawal struct {
	Index          string `json:"index"`
	ValidatorIndex string `json:"validatorIndex"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
}

type GetBlockNumberResp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type RequestBody struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}
