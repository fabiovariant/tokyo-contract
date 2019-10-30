package sql

import (
	"fmt"
	"database/sql"
	t "testing"

	"github.com/fabiovariant/tokyo-domains/contract"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewContract(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Client_Contract").
		WithArgs(c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, 
			c.DtInitContract, c.DtEndContract).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	
	repo := NewClientContractsSQLRepository(db)

	if err := repo.NewContract(&c); err != nil {
		t.Error(err)		
	}

	if c.ID != 1 {
		t.Error("Error on generated ID")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestNewContractOnFail(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO Client_Contract").
		WithArgs(c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, 
			c.DtInitContract, c.DtEndContract).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	
	repo := NewClientContractsSQLRepository(db)
	
	if err := repo.NewContract(&c); err == nil {
		t.Error(err)		
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetContractByClientID(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Client_Contract WHERE").
		WithArgs(c.ID).
		WillReturnRows(getContractRowsReturn(c))
	mock.ExpectCommit()

	repo := NewClientContractsSQLRepository(db)

	if cr, err := repo.GetContractByClientID(c.ID); err != nil || 
		cr.ID != c.ID {
		t.Error(err)
		t.Errorf("Returned ID: %b. Expected ID: %b", cr.ID, c.ID)	
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetContractByClientIDFailt(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT (.+) FROM Client_Contract WHERE").
		WithArgs(c.ID).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	
	repo := NewClientContractsSQLRepository(db)

	if _, err := repo.GetContractByClientID(c.ID); err == nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdate(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE Client_Contract").
		WithArgs(c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, 
			c.DtInitContract, c.DtEndContract, c.IsActive, c.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	
	repo := NewClientContractsSQLRepository(db)

	if err := repo.Update(&c); err != nil {
		t.Error(err)		
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}	
}

func TestUpdateFail(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	c := contract.GetContractMock()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE Client_Contract").
		WithArgs(c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, 
			c.DtInitContract, c.DtEndContract, c.IsActive, c.ID).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()
	
	repo := NewClientContractsSQLRepository(db)

	if err := repo.Update(&c); err == nil {
		t.Error(err)		
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}	
}

func getDbAndMock(t *t.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating Mock.")
	}
	return db, mock
}

func getContractRowsReturn(c contract.Contract) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"Client_Id", "Trade", "Company_Name", "Document_ID", "cd_Contract", "dt_init_contract", "dt_end_contract", "is_active"}).
		AddRow(c.ID, c.Trade, c.CompanyName, c.DocumentID, c.TpContract.CdContract, c.DtInitContract, c.DtEndContract, c.IsActive)
}