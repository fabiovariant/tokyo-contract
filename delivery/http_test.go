package delivery

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	t "testing"
	"time"

	"github.com/gorilla/mux"
	m "github.com/stretchr/testify/mock"

	"github.com/fabiovariant/tokyo-contracts/service"
	"github.com/fabiovariant/tokyo-domains/contract"
)

func TestNewContract(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()
	mock.On("NewContract", m.Anything).Return(nil)

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("POST", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.NewContract)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestNewContractServiceFail(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()
	mock.On("NewContract", m.Anything).Return(errors.New("Error"))

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("POST", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.NewContract)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestNewContractBadRequest(t *t.T) {
	mock := new(service.Mock)
	dm := NewClientContractHTTPDelivery(mock)

	// Request's mock object
	req, err := http.NewRequest("POST", "/house", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.NewContract)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertNotCalled(t, "NewContract")
}

func TestGetContractByClientID(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()
	mock.On("GetContractByClientID", m.AnythingOfType("int64")).Return(&contract, nil)

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("GET", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.GetContractByClientID)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestGetContractByClientIDBadRequest(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("GET", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "asdasd",
	})
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.GetContractByClientID)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusBadRequest)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertNotCalled(t, "GetContractByClientID")
}

func TestGetContractByClientIDServiceError(t *t.T) {
	contract := getValidContract()
	mock := new(service.Mock)
	mock.On("GetContractByClientID", m.AnythingOfType("int64")).Return(nil, errors.New("Error"))

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("GET", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.GetContractByClientID)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestUpdate(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()
	mock.On("Update", m.Anything).Return(nil)

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("PUT", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.Update)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestUpdateServiceFail(t *t.T) {
	mock := new(service.Mock)
	contract := getValidContract()
	mock.On("Update", m.Anything).Return(errors.New("Error"))

	dm := NewClientContractHTTPDelivery(mock)

	// Request's JSON
	b, _ := json.MarshalIndent(contract, "", "    ")
	// Request's mock object
	req, err := http.NewRequest("PUT", "/contract", strings.NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.Update)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertExpectations(t)
}

func TestUpdateBadRequest(t *t.T) {
	mock := new(service.Mock)
	dm := NewClientContractHTTPDelivery(mock)

	// Request's mock object
	req, err := http.NewRequest("PUT", "/house", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(dm.Update)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusInternalServerError)
		t.Errorf("Response body %v", rr.Body)
	}
	mock.AssertNotCalled(t, "Update")
}

func getValidContract() contract.Contract {
	return contract.Contract{
		Trade:       "Papada",
		CompanyName: "Tokyo Party",
		DocumentID:  "123123123",
		TpContract: contract.ContractType{
			CdContract: "Basic",
		},
		DtInitContract: time.Now().AddDate(-1, 0, 0),
		DtEndContract:  time.Now(),
		IsActive:       true,
	}
}
