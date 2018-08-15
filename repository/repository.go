package repository

import (
	"tokyo-house/domain"
)

// ContractsRepository is a interface to interact with the contracts layer on
// repository
type ContractsRepository interface {
	// NewContract add a new contract to repository
	NewContract(h *domain.House) (err error)

	// AllContracts list all contracts on repository
	AllContracts() (hs []*domain.House, err error)

	// GetContractByHouseID list a contract by house ID
	GetContractByHouseID(id int64) (h *domain.House, err error)

	// Update update a contract on repository
	Update(h *domain.House) (err error)
}
