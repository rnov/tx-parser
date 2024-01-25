package data

type CurrentBlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int    `json:"id"`
}

type BlockDataResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		BaseFeePerGas   string    `json:"baseFeePerGas"`
		Difficulty      string    `json:"difficulty"`
		ExtraData       string    `json:"extraData"`
		GasLimit        string    `json:"gasLimit"`
		GasUsed         string    `json:"gasUsed"`
		Hash            string    `json:"hash"`
		LogsBloom       string    `json:"logsBloom"`
		Miner           string    `json:"miner"`
		MixHash         string    `json:"mixHash"`
		Nonce           string    `json:"nonce"`
		Number          string    `json:"number"`
		ParentHash      string    `json:"parentHash"`
		ReceiptsRoot    string    `json:"receiptsRoot"`
		Sha3Uncles      string    `json:"sha3Uncles"`
		Size            string    `json:"size"`
		StateRoot       string    `json:"stateRoot"`
		Timestamp       string    `json:"timestamp"`
		TotalDifficulty string    `json:"totalDifficulty"`
		Transactions    []BlockTx `json:"transactions"`
	} `json:"result"`
}

// BlockTx holds the data of a transaction, some fields are omitted in order to make the response more readable.
type BlockTx struct {
	//BlockHash            string        `json:"blockHash"`
	//BlockNumber          string        `json:"blockNumber"`
	From     string `json:"from"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	//MaxFeePerGas         string        `json:"maxFeePerGas"`
	//MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	//Hash                 string        `json:"hash"`
	//Input                string        `json:"input"`
	//Nonce                string        `json:"nonce"`
	To string `json:"to"`
	//TransactionIndex     string        `json:"transactionIndex"`
	Value string `json:"value"`
	Type  string `json:"type"`
	//AccessList           []interface{} `json:"accessList"`
	ChainID string `json:"chainId"`
	//V                    string        `json:"v"`
	//R                    string        `json:"r"`
	//S                    string        `json:"s"`
	//YParity              string        `json:"yParity"`
}
