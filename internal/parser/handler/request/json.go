package request

import "tx-parser/pkg/data"

type AddAddressReq struct {
	Address string `json:"address"`
}

type GetBlockResp struct {
	Decimal int `json:"decimal"`
}

type GetTxsResp struct {
	Transactions []data.BlockTx `json:"transactions"`
}
