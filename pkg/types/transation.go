package types

type Transaction struct {
	ChainID     string `json:"chainId"`
	BlockNumber string `json:"blockNumber"`
	Hash        string `json:"hash"`
	Nonce       string `json:"nonce"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	Input       string `json:"input"`
}
