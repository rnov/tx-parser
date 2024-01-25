package virtual

import (
	"errors"
	"sync"
	"tx-parser/pkg/data"
)

// AddressHistoryVirtualStorage is an in-memory storage for address-txBlock data.
// It uses a read-write mutex for concurrent access protection.
type AddressHistoryVirtualStorage struct {
	// Embedding a pointer to a sync.RWMutex to protect concurrent access.
	*sync.RWMutex
	Storage map[string][]data.BlockTx
}

// NewVirtualStorage initializes and returns a new instance of AddressHistoryVirtualStorage.
// It sets up the internal map to store address data.
func NewVirtualStorage() *AddressHistoryVirtualStorage {
	return &AddressHistoryVirtualStorage{
		RWMutex: new(sync.RWMutex),
		Storage: make(map[string][]data.BlockTx),
	}
}

func (a *AddressHistoryVirtualStorage) AddAddress(address string) error {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	if _, ok := a.Storage[address]; !ok {
		a.Storage[address] = make([]data.BlockTx, 0)
	}
	return nil
}

func (a *AddressHistoryVirtualStorage) GetTransactions(address string) ([]data.BlockTx, error) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	if _, ok := a.Storage[address]; !ok {
		return nil, errors.New("address not found")
	}
	return a.Storage[address], nil
}

func (a *AddressHistoryVirtualStorage) GetAddresses() ([]string, error) {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	keys := make([]string, len(a.Storage))
	i := 0
	for k := range a.Storage {
		keys[i] = k
	}
	return keys, nil
}

func (a *AddressHistoryVirtualStorage) UpdateTxs(txs map[string][]data.BlockTx) error {
	a.RWMutex.Lock()
	defer a.RWMutex.Unlock()
	for k, v := range txs {
		// append new tx to existing txs
		a.Storage[k] = append(a.Storage[k], v...)
	}
	return nil
}
