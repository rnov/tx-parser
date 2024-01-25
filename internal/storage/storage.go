package storage

import (
	"tx-parser/internal/storage/virtual"
	"tx-parser/pkg/data"
)

// AddressHistory is an interface for address:txBlock storage.
type AddressHistory interface {
	AddAddress(address string) error
	GetAddresses() ([]string, error)
	GetTransactions(address string) ([]data.BlockTx, error)
	UpdateTxs(map[string][]data.BlockTx) error
}

// InitStorage initializes and returns a new instance of AddressHistory. based on the provided storage type.
func InitStorage(storageType string) AddressHistory {
	switch storageType {
	case "virtual":
		return virtual.NewVirtualStorage()
	//case "redis":
	//	return NewRedisStorage()
	default:
		return virtual.NewVirtualStorage()
	}
}
