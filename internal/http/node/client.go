package node

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	NodeUrl string
}

type NodeClient interface {
	GetCurrentBlockNumber() ([]byte, error)
	GetBlockData(blockNumber string) ([]byte, error)
}

// JSONRPCRequest represents a JSON-RPC request
type JSONRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// GetCurrentBlockNumber returns the current (latest) block number.
func (c *Client) GetCurrentBlockNumber() ([]byte, error) {
	requestBody, err := json.Marshal(JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []any{},
		ID:      1,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.NodeUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) GetBlockData(blockNumber string) ([]byte, error) {
	requestBody, err := json.Marshal(JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []any{blockNumber, true}, // true to get the full transaction objects
		ID:      1,
	})
	if err != nil {
		log.Fatalf("Error marshaling request: %v", err)
	}

	resp, err := http.Post(c.NodeUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error sending request for block %s: %v", blockNumber, err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	resp.Body.Close()

	return body, nil
}
