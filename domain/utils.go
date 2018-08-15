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

// GetMock returns a Mock of a House type to use in the test.
func GetMock() (house House) {
	house.Trade = "Tokio Party"
	house.CompanyName = "Tokio Party"
	house.DocumentID = "12.378.725/0001-05"
	house.Address = Address{
		Address: "Rua Augusta",
		Number:  "666",
		State:   "São Paulo",
		City:    "São Paulo",
		Region:  "Centro"}
	house.Styles = []*Style{&Style{CdStyle: "POP", NmStyle: "Pop", DsStyle: "Estilo dançante"}}
	house.Desc = "Casa com decoração retro"
	house.HasSmokingArea = true
	house.NumberOfFloors = 3

	house.Location = GeoJSON{
		Type:        "Point",
		Coordinates: [2]float64{-23.5521025, -46.6543741}}

	house.Products = []*Product{
		&Product{
			Product: "Cerveja Skol",
			Price:   "R$: 9,00",
			SubType: ProductSubType{
				CdSubType: "NACIO",
				NmSubType: "Cervejas nacionais",
				Type: ProductType{
					CdType: "CERV",
					NmProd: "Cervejas",
				},
			},
		}}

	house.Entertainments = []*Entertainment{&Entertainment{
		Name:            "Mesa de sinuca",
		CdEntertainment: "jogos"}}

	house.Contract = HouseContract{
		CdContract:     "CT_01",
		DsContract:     "Contrato padrão",
		DtInitContract: time.Now(),
		DtEndContract:  time.Now().AddDate(2, 0, 0),
	}
	return
}
