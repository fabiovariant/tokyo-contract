package sql

import (
	"database/sql"

	"github.com/fabiovariant/tokyo-contracts/repository"
	"github.com/fabiovariant/tokyo-domains/contract"
)

type clientContractsSQLRepository struct {
	db *sql.DB
}

// NewClientContractsSQLRepository returns a new instance of finalCustomerRepository.
func NewClientContractsSQLRepository(db *sql.DB) repository.ClientContractsRepository {
	return &clientContractsSQLRepository{db}
}

func (cr *clientContractsSQLRepository) NewContract(c *contract.Contract) (err error) {
	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}
	defer cr.deferTxFunc()(tx, &err)

	sql := `INSERT INTO Client_Contract (Trade, Company_Name, Document_ID, cd_Contract, dt_init_contract, dt_end_contract) VALUES (?, ?, ?, ?, ?, ?)`
	rs, err := tx.Exec(sql, c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, c.DtInitContract, c.DtEndContract)
	if err != nil {
		return err
	}
	c.ID, err = rs.LastInsertId()
	return
}

func (cr *clientContractsSQLRepository) GetContractByClientID(id int64) (c *contract.Contract, err error) {
	tx, err := cr.db.Begin()
	if err != nil {
		return nil, err
	}
	defer cr.deferTxFunc()(tx, &err)

	sql := `SELECT Client_Id, Trade, Company_Name, Document_ID, cd_Contract, dt_init_contract, dt_end_contract, is_active
				FROM Client_Contract WHERE Client_Id = $1`
	rows, err := tx.Query(sql, id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		var ct contract.ContractType
		var cs contract.Contract

		err = rows.Scan(&cs.ID, &cs.Trade, &cs.CompanyName, &cs.DocumentID, &ct.CdContract,
			&cs.DtInitContract, &cs.DtEndContract, &cs.IsActive)

		c = &cs
		c.TpContract = ct
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (cr *clientContractsSQLRepository) Update(c *contract.Contract) (err error) {
	return
}

func (cr *clientContractsSQLRepository) deferTxFunc() func(tx *sql.Tx, err *error) {

	return func(tx *sql.Tx, err *error) {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			if er := tx.Commit(); er != nil {
				panic(er)
			}
		}
	}
}
