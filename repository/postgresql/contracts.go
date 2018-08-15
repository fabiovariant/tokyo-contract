package postgresql

import (
	"database/sql"
	"tokyo-house/domain"
	"tokyo-house/repository"
)

type contractsPostgresqlRepository struct {
	Db *sql.DB
}

// NewContractsPostgresqlRepository returns a new instance of finalCustomerRepository.
func NewContractsPostgresqlRepository(db *sql.DB) repository.ContractsRepository {
	return &contractsPostgresqlRepository{db}
}

func (cr *contractsPostgresqlRepository) NewContract(h *domain.House) (err error) {
	sql := `INSERT INTO House_Contract
				(Trade,
				Company_Name,
				Document_ID,
				cd_Contract,
				dt_init_contract,
				dt_end_contract,
				is_active)
			VALUES (?, ?, ?, ?, ?, ?, ?)`
	tx := getConnTX(cr.Db)
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	r, err := tx.Exec(sql, h.Trade, h.CompanyName, h.DocumentID, h.Contract.CdContract,
		h.Contract.DtInitContract, h.Contract.DtEndContract, h.Contract.IsActive)
	if err != nil {
		return
	}
	h.ID, err = r.LastInsertId()
	return
}

func (cr *contractsPostgresqlRepository) AllContracts() (hs []*domain.House, err error) {
	sql := `SELECT
				House_id,
				Trade,
				Company_Name,
				Document_ID,
				cd_Contract,
				dt_init_contract,
				dt_end_contract,
				is_active
			FROM House_Contract`
	rows, err := cr.Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		var h = new(domain.House)
		err = rows.Scan(&h.ID, &h.Trade, &h.CompanyName, &h.DocumentID, &h.Contract.CdContract,
			&h.Contract.DtInitContract, &h.Contract.DtEndContract, &h.Contract.IsActive)
		if err != nil {
			return
		}
		hs = append(hs, h)
	}
	return
}

func (cr *contractsPostgresqlRepository) GetContractByHouseID(id int64) (h *domain.House, err error) {
	sql := `SELECT
				House_id,
				Trade,
				Company_Name,
				Document_ID,
				cd_Contract,
				dt_init_contract,
				dt_end_contract,
				is_active
			FROM House_Contract WHERE House_id = ?`
	rows, err := cr.Db.Query(sql, id)
	if err != nil {
		return
	}
	if rows.Next() {
		hh := new(domain.House)
		err = rows.Scan(&hh.ID, &hh.Trade, &hh.CompanyName, &hh.DocumentID, &hh.Contract.CdContract,
			&hh.Contract.DtInitContract, &hh.Contract.DtEndContract, &hh.Contract.IsActive)
		if err != nil {
			return
		}
		h = hh
	}
	return
}

func (cr *contractsPostgresqlRepository) Update(h *domain.House) (err error) {
	sql := `UPDATE House_Contract SET
				Trade = $1,
				Company_Name = $2,
				Document_ID = $3,
				cd_Contract = $4,
				dt_init_contract = $5,
				dt_end_contract = $6,
				is_active = $7
			WHERE House_id = $8`
	tx := getConnTX(cr.Db)
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(sql, h.Trade, h.CompanyName, h.DocumentID, h.Contract.CdContract,
		h.Contract.DtInitContract, h.Contract.DtEndContract, h.Contract.IsActive, h.ID)
	return
}

func getConnTX(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	if err != nil {
		panic("Não foi possível iniciar transação.")
	}
	return tx
}