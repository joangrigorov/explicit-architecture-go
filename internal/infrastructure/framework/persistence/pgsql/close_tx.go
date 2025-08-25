package pgsql

import "database/sql/driver"

func CloseTx(tx driver.Tx, err *error) {
	if r := recover(); r != nil {
		_ = tx.Rollback()
		panic(r)
	} else if *err != nil {
		_ = tx.Rollback()
	} else {
		*err = tx.Commit()
	}
}
