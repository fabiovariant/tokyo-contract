package delivery

import (
	"net/http"
)

// ClientContractsDelivery is a interface to interact with the contracts layer on
// repository
type ClientContractsDelivery interface {
	// NewContract add a new contract to repository
	NewContract(w http.ResponseWriter, r *http.Request)

	// GetContractByClientID list a contract by house ID
	GetContractByClientID(w http.ResponseWriter, r *http.Request)

	// Update update a contract on repository
	Update(w http.ResponseWriter, r *http.Request)
}
