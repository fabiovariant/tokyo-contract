package service

import (
	"github.com/fabiovariant/tokyo-domains/contract"
)

// ClientContractsService is a interface to interact with the contracts layer on
// repository
type ClientContractsService interface {
	// NewContract add a new contract to repository
	NewContract(h *contract.Contract) (err error)

	// GetContractByClientID list a contract by house ID
	GetContractByClientID(id int64) (h *contract.Contract, err error)

	// Update update a contract on repository
	Update(h *contract.Contract) (err error)
}
