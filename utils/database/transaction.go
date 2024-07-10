package database

import (
	"cake-store/utils/wrapper"
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		wrapper.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		wrapper.PanicIfError(errorCommit)
	}
}
