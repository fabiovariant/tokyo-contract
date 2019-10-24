package service

import (
	"github.com/fabiovariant/tokyo-contracts/repository"
	"github.com/fabiovariant/tokyo-domains/contract"
)

type clientContractsService struct {
	repository repository.ClientContractsRepository
}

// NewClientContractService returns a instantiate service type
func NewClientContractService(r repository.ClientContractsRepository) ClientContractsService {
	return &clientContractsService{r}
}

func (cs *clientContractsService) NewContract(c *contract.Contract) (err error) {
	return cs.repository.NewContract(c)
}

func (cs *clientContractsService) GetContractByClientID(id int64) (c *contract.Contract, err error) {
	return cs.repository.GetContractByClientID(id)
}

func (cs *clientContractsService) Update(c *contract.Contract) (err error) {
	return
}
