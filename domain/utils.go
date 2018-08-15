package domain

import (
	"encoding/json"
	"net/http"
	"time"
)

//WriteJSONResponse write a JSON response of the House domain.
func WriteJSONResponse(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

// GetMock returns a Mock of a HouseContract type to use in the test.
func GetMock() (ctt HouseContract) {
	ctt.ID = 1
	ctt.Trade = "Tokyo Party"
	ctt.CompanyName = "Tokyo Party"
	ctt.DocumentID = "1303108760"
	ctt.DtInitContract = time.Now()
	ctt.DtEndContract = time.Now()
	ctt.IsActive = true
	ctt.TpContract = ContractType{CdContract: "CT01",
		DsContract: "Contrato padrão com rescursos básicos."}
	return
}
