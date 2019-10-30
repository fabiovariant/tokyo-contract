package repository

import (
	"github.com/fabiovariant/tokyo-domains/contract"

	"github.com/stretchr/testify/mock"
)

// Mock type of repo
type Mock struct {
	mock.Mock
}

// NewContract mock to repository layer.
func (mock *Mock) NewContract(h *contract.Contract) (err error) {
	args := mock.Called(h)
	err = args.Error(0)
	return
}

// GetContractByClientID mock to repository layer.
func (mock *Mock) GetContractByClientID(id int64) (c *contract.Contract, err error) {
	args := mock.Called(id)

	if args.Get(0) != nil {
		c = args.Get(0).(*contract.Contract)
	}
	err = args.Error(1)
	return
}

// Update mock to repository layer.
func (mock *Mock) Update(h *contract.Contract) (err error) {
	args := mock.Called(h)
	err = args.Error(0)
	return
}
