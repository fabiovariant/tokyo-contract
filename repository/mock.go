package repository

import (
	"tokyo-house/domain"

	"github.com/stretchr/testify/mock"
)

// Mock type of repo
type Mock struct {
	mock.Mock
}

// NewContract mock to repository layer.
func (mock *Mock) NewContract(h *domain.House) (err error) {
	return
}

// AllContracts mock to repository layer.
func (mock *Mock) AllContracts() (hs []*domain.House, err error) {
	return
}

// GetContractByHouseID mock to repository layer.
func (mock *Mock) GetContractByHouseID(id int64) (h *domain.House, err error) {
	return
}

// Update mock to repository layer.
func (mock *Mock) Update(h *domain.House) (err error) {
	return
}
