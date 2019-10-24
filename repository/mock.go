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
	return
}

// GetContractByClientID mock to repository layer.
func (mock *Mock) GetContractByClientID(id int64) (h *contract.Contract, err error) {
	return
}

// Update mock to repository layer.
func (mock *Mock) Update(h *contract.Contract) (err error) {
	return
}
