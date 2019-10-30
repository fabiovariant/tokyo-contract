package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/fabiovariant/tokyo-contracts/service"
	"github.com/fabiovariant/tokyo-domains/contract"
	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("delivery-http")

// ClientContractsHTTPDelivery http delivery impl
type ClientContractsHTTPDelivery struct {
	s service.ClientContractsService
}

// NewClientContractHTTPDelivery returns a instantiate service type
func NewClientContractHTTPDelivery(s service.ClientContractsService) ClientContractsDelivery {
	return &ClientContractsHTTPDelivery{s}
}

// NewContract http impl
func (cht *ClientContractsHTTPDelivery) NewContract(w http.ResponseWriter, r *http.Request) {
	var c contract.Contract
	if err := jsonToContract(r.Body, &c); err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal Bad Request", http.StatusBadRequest)
		return
	}

	err := cht.s.NewContract(&c)
	if err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	simpleMessageWithJSON(w, "Success", http.StatusOK)
}

// GetContractByClientID http impl
func (cht *ClientContractsHTTPDelivery) GetContractByClientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	q := vars["id"]
	contractID, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal Bad Request", http.StatusBadRequest)
		return
	}

	c, err := cht.s.GetContractByClientID(contractID)
	if err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	simpleMessageWithJSON(w, c, http.StatusOK)
}

// Update http impl
func (cht *ClientContractsHTTPDelivery) Update(w http.ResponseWriter, r *http.Request) {
	var c contract.Contract
	if err := jsonToContract(r.Body, &c); err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal Bad Request", http.StatusBadRequest)
		return
	}

	err := cht.s.Update(&c)
	if err != nil {
		log.Error("Error: ", err)
		simpleMessageWithJSON(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	simpleMessageWithJSON(w, "Success", http.StatusOK)
}

func jsonToContract(r io.Reader, c *contract.Contract) error {
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return errors.New("Erro ao converter JSON \n" + err.Error())
	}
	return nil
}

func simpleMessageWithJSON(w http.ResponseWriter, message interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
