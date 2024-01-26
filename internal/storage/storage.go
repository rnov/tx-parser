package storage

import (
	"tx-parser/internal/storage/virtual"
	"tx-parser/pkg/data"
)

// AddressHistory is an interface for address:txBlock storage.
// The structure of the methods is being though to be used with different storage types.
// that's the reason to add error returns to all methods, although they are not used in the current implementation.
type AddressHistory interface {
	AddAddress(address string) error
	CheckAddress(address string) (bool, error)
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
