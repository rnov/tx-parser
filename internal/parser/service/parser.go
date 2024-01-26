package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"tx-parser/internal/http/node"
	"tx-parser/internal/storage"
	"tx-parser/pkg/data"
)

type blockChecker struct {
	hexCurrentBlock string
	decCurrentBlock int64
	intervalCheck   int
}

// Service implements Parser (service) interface methods and holds all the resources that are need to carry its functionality.
type Service struct {
	nc     node.NodeClient
	blkChk blockChecker
	store  storage.AddressHistory
}

func NewService(nodeURL, storageType string, checkTime int) *Service {
	nc := &node.Client{NodeUrl: nodeURL}
	store := storage.InitStorage(storageType)
	bc := blockChecker{intervalCheck: checkTime}

	return &Service{nc: nc, blkChk: bc, store: store}
}

// StartChecker initiates in parallel continuous checks for the latest block number updating the block and triggering check on
// block data for transaction involving the subscribed addresses.
func (s *Service) StartChecker() {
	// get latest block number
	go func() {
		fmt.Println("node checker initiated...")
		for {
			res, err := s.nc.GetCurrentBlockNumber()
			if err != nil {
				fmt.Printf("error retrieving block from node: %s\n", err.Error())
			}
			hexV, decV, err := parseBlockValue(res)
			// valid response and new block -> check block data for transactions involving our addresses
			if decV != 0 && s.blkChk.decCurrentBlock < decV {
				s.blkChk.hexCurrentBlock = hexV
				s.blkChk.decCurrentBlock = decV
				// note: could have avoided triggering a goroutine as well as we could also use a chan to send the bytes and do
				// data retrieval and check for tx in another thread. leaved this way for simplicity.
				go s.getBlockData()
			}
			time.Sleep(time.Duration(s.blkChk.intervalCheck) * time.Second)
		}
	}()

}

// Parser interface exposes its functionalities
type Parser interface {
	// last parsed block
	GetCurrentBlock() int

	// add address to observer
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []data.BlockTx
}

func (s *Service) GetCurrentBlock() int {
	// note: for simplicity i've not added any mutex neither added the block in the storage.
	return int(s.blkChk.decCurrentBlock)
}

func (s *Service) Subscribe(address string) bool {
	// note: as noted in the description, simplicity, not external libs, not production ready: therefore not checking the address structure
	if exist, _ := s.store.CheckAddress(address); !exist {
		if err := s.store.AddAddress(address); err != nil {
			fmt.Printf("error adding address to store: %s\n", err.Error())
			return false
		}
	}

	return true
}

func (s *Service) GetTransactions(address string) []data.BlockTx {
	res, err := s.store.GetTransactions(address)
	if err != nil {
		fmt.Printf("error getting txs from store: %s\n", err.Error())
		return make([]data.BlockTx, 0)
	}
	return res
}

// parseBlockValue parses block value to hexadecimal and decimal value.
func parseBlockValue(b []byte) (string, int64, error) {
	lastBlock := &data.CurrentBlockResponse{}
	err := json.Unmarshal(b, lastBlock)
	if err != nil {
		fmt.Printf("error retrieving last block response data: %s\n", err.Error())
		return "", 0, err
	}
	hexValue := lastBlock.Result
	decimalValue, err := strconv.ParseInt(lastBlock.Result, 0, 64)
	if err != nil {
		fmt.Println("Error converting hex to int:", err)
		return "", 0, err
	}
	return hexValue, decimalValue, err
}

// getBlockData retrieves block data(txs) from the node and looks for txs involving addresses that have been subscribed
func (s *Service) getBlockData() {
	resp, err := s.nc.GetBlockData(s.blkChk.hexCurrentBlock)
	if err != nil {
		fmt.Printf("error retrieving data from node: %s\n", err.Error())
		return
	}

	var rpcResp data.BlockDataResponse
	if err := json.Unmarshal(resp, &rpcResp); err != nil {
		log.Fatalf("Error unmarshalling response: %v", err.Error())
	}

	// get all addresses from storage
	addrs, err := s.store.GetAddresses()
	if err != nil {
		log.Fatalf("Error retrieving data from storage: %v", err.Error())
	}

	//note: initialize an aux Map with all the addresses as keys, in order to iterate just once the result tx with linear complexity
	// and O(1) in accessing a hash rather than iterating the entire slice for every txs.
	auxAddrMap := make(map[string][]data.BlockTx)
	for _, addr := range addrs {
		auxAddrMap[addr] = make([]data.BlockTx, 0)
	}

	// iterate all txs checking fields 'from' and 'to' if any matches our addresses
	for _, tx := range rpcResp.Result.Transactions {
		//fmt.Printf("from: '%s' to: '%s'\n", tx.From, tx.To)
		addressHit := ""

		// check for tx addresses
		if _, ok := auxAddrMap[strings.ToLower(tx.From)]; ok {
			addressHit = strings.ToLower(tx.From)
			// add tx to the aux map if and address has been found
			if addressHit != "" {
				auxAddrMap[addressHit] = append(auxAddrMap[addressHit], tx)
			}
			addressHit = ""
		}

		if _, ok := auxAddrMap[strings.ToLower(tx.To)]; ok {
			addressHit = strings.ToLower(tx.To)
			// add tx to the aux map if and address has been found
			if addressHit != "" {
				auxAddrMap[addressHit] = append(auxAddrMap[addressHit], tx)
			}
		}
	}
	// update address storage
	if err := s.store.UpdateTxs(auxAddrMap); err != nil {
		fmt.Printf("error updating data storage: %s\n", err.Error())
	}
}
