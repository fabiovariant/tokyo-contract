package service

import (
	"errors"
	
	"github.com/fabiovariant/tokyo-domains/contract"
	"github.com/fabiovariant/tokyo-contracts/repository"
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

type clientContractsService struct {
	repository repository.ClientContractsRepository
}

// NewClientContractService returns a instantiate service type
func NewClientContractService(r repository.ClientContractsRepository) ClientContractsService {
	return &clientContractsService{r}
}

func (cs *clientContractsService) NewContract(c *contract.Contract) (err error) {
	if err = c.ValidateData(); err != nil {
		return err
	}
	return cs.repository.NewContract(c)
}

func (cs *clientContractsService) GetContractByClientID(id int64) (c *contract.Contract, err error) {
	return cs.repository.GetContractByClientID(id)
}

func (cs *clientContractsService) Update(c *contract.Contract) (err error) {
	if err = c.ValidateData(); err != nil {
		return err
	}
	if c.ID == 0 {
		return errors.New("Invalid ID")
	}
	return cs.repository.Update(c)
}
