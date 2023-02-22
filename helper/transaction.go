package helper

import "database/sql"

func CommitRollback(transaction *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := transaction.Rollback()
		PanicError(errorRollback)
		panic(err)
	} else {
		errorCommit := transaction.Commit()
		PanicError(errorCommit)
	}
}
