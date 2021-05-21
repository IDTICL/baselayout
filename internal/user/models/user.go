package models

import "database/sql"

type UserModel struct {
	DB      *sql.DB
	Read    *sql.DB
	Write   *sql.DB
	TxRead  *sql.Tx
	TxWrite *sql.Tx
}

func (um *UserModel) Insert() error {
	return nil
}
