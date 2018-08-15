package domain

import (
	"encoding/json"
	"time"
)

// HouseContract type to reflect the contracts layer.
type HouseContract struct {
	ID             int64        `json:"id"`
	Trade          string       `json:"trade"`
	CompanyName    string       `json:"companyName"`
	DocumentID     string       `json:"documentId"`
	TpContract     ContractType `json:"cd_contract"`
	DtInitContract time.Time    `json:"dt_init_contract"`
	DtEndContract  time.Time    `json:"dt_end_contract"`
	IsActive       bool         `json:"is_Active"`
}

func (hc *HouseContract) String() string {
	b, _ := json.MarshalIndent(hc, "", "    ")
	return string(b)
}

// ContractType reflects the contract type database structure.
type ContractType struct {
	CdContract string `json:"cd_Contract"`
	DsContract string `json:"ds_contract"`
}

func (ctt *ContractType) String() string {
	b, _ := json.MarshalIndent(ctt, "", "    ")
	return string(b)
}
