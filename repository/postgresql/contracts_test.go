package postgresql

import (
	"database/sql"
	"errors"
	t "testing"
	"tokyo-house/domain"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewContract(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO House_Contract").
		WithArgs(house.Trade, house.CompanyName, house.DocumentID, house.Contract.CdContract,
			house.Contract.DtInitContract, house.Contract.DtEndContract, house.Contract.IsActive).
		WillReturnResult(sqlmock.NewResult(house.ID, 1))
	mock.ExpectCommit()
	a := NewContractsPostgresqlRepository(db)
	if err := a.NewContract(&house); err != nil {
		t.Errorf("Error saving new house.")
		t.Error(err)
	}
	if house.ID != house.ID {
		t.Error("Error on returned ID.")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestNewContractOnFailure(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO House_Contract").
		WithArgs(house.Trade, house.CompanyName, house.DocumentID, house.Contract.CdContract,
			house.Contract.DtInitContract, house.Contract.DtEndContract, house.Contract.IsActive).
		WillReturnError(errors.New("Error"))
	mock.ExpectRollback()
	a := NewContractsPostgresqlRepository(db)
	if err := a.NewContract(&house); err == nil {
		t.Errorf("Error saving new house.")
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestAllContracts(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	rows := sqlmock.NewRows([]string{
		"House_id",
		"Trade",
		"Company_Name",
		"Document_ID",
		"cd_Contract",
		"dt_init_contract",
		"dt_end_contract",
		"is_active"}).
		AddRow(
			house.ID,
			house.Trade,
			house.CompanyName,
			house.DocumentID,
			house.Contract.CdContract,
			house.Contract.DtInitContract,
			house.Contract.DtEndContract,
			house.Contract.IsActive)

	mock.ExpectQuery("SELECT (.+) FROM House_Contract").WillReturnRows(rows)

	a := NewContractsPostgresqlRepository(db)
	hs, err := a.AllContracts()
	if err != nil {
		t.Errorf("Error getting all house.")
		t.Error(err)
	}
	if hs[0].ID != house.ID {
		t.Errorf("Error on returned house.")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestAllContractsOnFailure(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	mock.ExpectQuery("SELECT (.+) FROM House_Contract").
		WillReturnError(errors.New("Error"))

	a := NewContractsPostgresqlRepository(db)
	if _, err := a.AllContracts(); err == nil {
		t.Errorf("Error getting all house contracts failure test.")
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestGetContractByHouseID(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	rows := sqlmock.NewRows([]string{
		"House_id",
		"Trade",
		"Company_Name",
		"Document_ID",
		"cd_Contract",
		"dt_init_contract",
		"dt_end_contract",
		"is_active"}).
		AddRow(
			house.ID,
			house.Trade,
			house.CompanyName,
			house.DocumentID,
			house.Contract.CdContract,
			house.Contract.DtInitContract,
			house.Contract.DtEndContract,
			house.Contract.IsActive)

	mock.ExpectQuery("SELECT (.+) FROM House_Contract WHERE House_id(.*)").WillReturnRows(rows)

	a := NewContractsPostgresqlRepository(db)
	hr, err := a.GetContractByHouseID(house.ID)
	if err != nil {
		t.Errorf("Error getting all house contract.")
		t.Error(err)
	}
	if hr.ID != house.ID {
		t.Errorf("Error on returned house contract.")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestGetContractByHouseIDOnFailure(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	mock.ExpectQuery("SELECT (.+) FROM House_Contract WHERE House_id(.*)").
		WillReturnError(errors.New("Error"))

	a := NewContractsPostgresqlRepository(db)
	if _, err := a.GetContractByHouseID(house.ID); err == nil {
		t.Errorf("Error getting all house contracts failure test.")
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestUpdateContract(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE House_Contract SET").
		WithArgs(house.Trade, house.CompanyName, house.DocumentID, house.Contract.CdContract,
			house.Contract.DtInitContract, house.Contract.DtEndContract, house.Contract.IsActive, house.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	a := NewContractsPostgresqlRepository(db)
	if err := a.Update(&house); err != nil {
		t.Errorf("Error updating house contract.")
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func TestUpdateContractOnFailure(t *t.T) {
	db, mock := getDbAndMock(t)
	defer db.Close()
	house := domain.GetMock()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE House_Contract SET").
		WithArgs(house.Trade, house.CompanyName, house.DocumentID, house.Contract.CdContract,
			house.Contract.DtInitContract, house.Contract.DtEndContract, house.Contract.IsActive, house.ID).
		WillReturnError(errors.New("Error"))
	mock.ExpectRollback()
	a := NewContractsPostgresqlRepository(db)
	if err := a.Update(&house); err == nil {
		t.Errorf("Error updating house contract.")
		t.Error(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Something wrong: %s", err)
	}
}

func getDbAndMock(t *t.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating Mock.")
	}
	return db, mock
}