package domain

import (
	"encoding/json"
	"time"
)

// House is the house type
type House struct {
	ID             int64            `json:"id"`
	Trade          string           `json:"trade"`
	CompanyName    string           `json:"companyName"`
	DocumentID     string           `json:"documentId"`
	Email          string           `json:"email"`
	Desc           string           `json:"desc"`
	HasSmokingArea bool             `json:"hasSmokingArea"`
	NumberOfFloors int              `json:"numberOfFloors"`
	UserDistance   float64          `json:"userDistanceFrom"`
	Address        Address          `json:"address"`
	Location       GeoJSON          `json:"location"`
	Styles         []*Style         `json:"styles"`
	Products       []*Product       `json:"products"`
	Entertainments []*Entertainment `json:"entertainment"`
	Contract       HouseContract    `json:"house_contract"`
}

func (h House) String() string {
	b, _ := json.MarshalIndent(h, "", "    ")
	return string(b)
}

//Product is a house's product
type Product struct {
	Product     string         `json:"product"`
	Price       string         `json:"price"`
	CdProd      string         `json:"cd_prod"`
	Observation string         `json:"obs"`
	SubType     ProductSubType `json:"sub_type"`
}

func (p Product) String() string {
	b, _ := json.MarshalIndent(p, "", "    ")
	return string(b)
}

//ProductSubType reflects the sub type data structure
type ProductSubType struct {
	CdSubType string      `json:"cd_sub_type"`
	NmSubType string      `json:"nm_sub_type"`
	Type      ProductType `json:"product_type"`
}

func (pst ProductSubType) String() string {
	b, _ := json.MarshalIndent(pst, "", "    ")
	return string(b)
}

//ProductType reflects the prod type structure
type ProductType struct {
	CdType string `json:"cd_type"`
	NmProd string `json:"nm_prod"`
}

func (pt ProductType) String() string {
	b, _ := json.MarshalIndent(pt, "", "    ")
	return string(b)
}

// Style is a house's product
type Style struct {
	CdStyle string `json:"style"`
	NmStyle string `json:"nm_style"`
	DsStyle string `json:"ds_style"`
}

func (s Style) String() string {
	b, _ := json.MarshalIndent(s, "", "    ")
	return string(b)
}

//Entertainment that the house offers
type Entertainment struct {
	Name            string `json:"name"`
	CdEntertainment string `json:"cd_entertainment"`
}

func (e Entertainment) String() string {
	b, _ := json.MarshalIndent(e, "", "    ")
	return string(b)
}

//Address type to house address.
type Address struct {
	Address    string `json:"address"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	State      string `json:"state"`
	City       string `json:"city"`
	Region     string `json:"region"`
}

func (a Address) String() string {
	b, _ := json.MarshalIndent(a, "", "    ")
	return string(b)
}

// GeoJSON geospatial data
type GeoJSON struct {
	Type        string     `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

func (g GeoJSON) String() string {
	b, _ := json.MarshalIndent(g, "", "    ")
	return string(b)
}

// HouseContract type to reflect the contracts layer.
type HouseContract struct {
	CdContract     string    `json:"cd_contract"`
	DsContract     string    `json:"ds_contract"`
	DtInitContract time.Time `json:"dt_init_contract"`
	DtEndContract  time.Time `json:"dt_end_contract"`
	IsActive       bool      `json:"is_Active"`
}

func (hc *HouseContract) String() string {
	b, _ := json.MarshalIndent(hc, "", "    ")
	return string(b)
}
