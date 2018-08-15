BEGIN;

/*
    Domain register of contract types to reference.
*/
CREATE TABLE IF NOT EXISTS Contract_Types (
    cd_Contract             varchar(5)              NOT NULL,
    ds_contract             character varying,
    PRIMARY KEY(cd_Contract)
);

/*
    Register of house contracts.
    A contract must to be active and with dt_end_contract <= now()
*/
CREATE TABLE IF NOT EXISTS House_Contract (
    House_id                serial                  NOT NULL,
    Trade                   character varying       NOT NULL,
    Company_Name            character varying       NOT NULL,
    Document_ID             character varying       NOT NULL,
    cd_Contract             varchar(5)              NOT NULL,
    dt_init_contract        timestamp               NOT NULL DEFAULT NOW(),
    dt_end_contract         timestamp               NOT NULL,
    is_active               boolean                 NOT NULL,
    PRIMARY KEY(house_id),
    FOREIGN KEY(cd_Contract) REFERENCES Contract_Types(cd_Contract)
);

ROLLBACK;