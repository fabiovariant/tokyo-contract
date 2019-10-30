package service

import (
	"time"
	"errors"
	t "testing"

	"github.com/stretchr/testify/assert"
	"github.com/fabiovariant/tokyo-domains/contract"
	"github.com/fabiovariant/tokyo-contracts/repository"
)

func TestNewContract(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	repoMock.On("NewContract", &contract).Return(nil)
	service := NewClientContractService(repoMock)
	if err := service.NewContract(&contract); err != nil {
		t.Error("Error on new contract test")
		t.Error(err)
	}
	repoMock.AssertExpectations(t)
}

func TestNewContractRepositoryError(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	repoMock.On("NewContract", &contract).Return(errors.New("Error"))
	service := NewClientContractService(repoMock)
	if err := service.NewContract(&contract); err == nil {
		t.Error("Error on new contract fail test")
	}
	repoMock.AssertExpectations(t)
}

func TestNewContractDataError(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	contract.Trade = ""
	service := NewClientContractService(repoMock)
	if err := service.NewContract(&contract); err == nil {
		t.Error("Error on new contract fail test")
	}
	repoMock.AssertNotCalled(t, "NewContract", &contract)
}

func TestGetContractByClientID(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	repoMock.On("GetContractByClientID", contract.ID).Return(&contract, nil)
	service := NewClientContractService(repoMock)
	cr, err := service.GetContractByClientID(contract.ID)
	if err != nil {
		t.Error("Error on new contract test")
		t.Error(err)
	}

	assert.Equal(t, cr.ID, contract.ID, "Not equal")
	repoMock.AssertExpectations(t)	
}

func TestGetContractByClientIDRepoFail(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	repoMock.On("GetContractByClientID", contract.ID).Return(nil, errors.New("Error"))
	service := NewClientContractService(repoMock)
	cr, err := service.GetContractByClientID(contract.ID)
	if err == nil {
		t.Error("Error on new contract test")
	}

	assert.Nil(t, cr)
	repoMock.AssertExpectations(t)	
}

func TestUpdate(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	contract.ID = 10
	repoMock.On("Update", &contract).Return(nil)
	service := NewClientContractService(repoMock)
	if err := service.Update(&contract); err != nil {
		t.Error("Error on update contract test")
		t.Error(err)
	}
	repoMock.AssertExpectations(t)
}

func TestUpdateIDFail(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	service := NewClientContractService(repoMock)
	if err := service.Update(&contract); err == nil {
		t.Error("Error on update contract fail test")
	}
	repoMock.AssertNotCalled(t, "Update", &contract)
}

func TestUpdateDataFail(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	contract.ID = 10
	contract.Trade = ""
	service := NewClientContractService(repoMock)
	if err := service.Update(&contract); err == nil {
		t.Error("Error on update contract fail test")
	}
	repoMock.AssertNotCalled(t, "Update", &contract)
}

func TestUpdateRepoFail(t *t.T) {
	repoMock := new(repository.Mock)
	contract := getValidContract()
	contract.ID = 10
	repoMock.On("Update", &contract).Return(errors.New("Error"))
	service := NewClientContractService(repoMock)
	if err := service.Update(&contract); err == nil {
		t.Error("Error on update contract fail test")
	}
	repoMock.AssertExpectations(t)
}

func getValidContract() contract.Contract {
	return contract.Contract {
		Trade: "Papada",
		CompanyName: "Tokyo Party",
		DocumentID: "123123123",
		TpContract: contract.ContractType{
			CdContract: "Basic",
		},
		DtInitContract: time.Now().AddDate(-1, 0, 0),
		DtEndContract: time.Now(),
		IsActive: true,
	}
}