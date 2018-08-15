package repository

import (
	"tokyo-house-contract/domain"
)

// ContractsRepository is a interface to interact with the contracts layer on
// repository
type ContractsRepository interface {
	// NewContract add a new contract to repository
	NewContract(h *domain.HouseContract) (err error)

	// AllContracts list all contracts on repository
	AllContracts() (hs []*domain.HouseContract, err error)

	// GetContractByHouseID list a contract by house ID
	GetContractByHouseID(id int64) (h *domain.HouseContract, err error)

	// Update update a contract on repository
	Update(h *domain.HouseContract) (err error)
}
