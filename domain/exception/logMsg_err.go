package exception

const (
	LogErrDBOpenConn   = "ERROR OPEN CONNECTION DB"
	LogErrDBCloseConn  = "ERROR CLOSE CONNECTION DB FOR RETURNING TO POOL"
	LogErrDBCloseStmt  = "ERROR CLOSE PREPARED STATEMENT FOR FREE TO MEMORY"
	LogErrDBCloseRows  = "ERROR CLOSE ROWS FOR FREE FOR DATA"
	LogErrDBScanning   = "ERROR SCANNING DATA"
	LogErrDBStmt       = "ERROR START PREPARED STATEMENT"
	LogErrDBExec       = "ERROR EXEC TO DATA"
	LogErrDBQuery      = "ERROR QUERY TO DATA"
	LogErrDBStartTx    = "ERROR START TRANSACTION"
	LogErrDBTx         = "ERROR TRANSACTION"
	LogErrDBTxNil      = "ERROR TRANSACTION IS NIL"
	LogErrDBTxRollback = "ERROR ROLLBACK TRANSACTION"
	LogErrDBTxCommit   = "ERROR COMMIT TRANSACTION"
	LogErrDBTxStart    = "CANNOT START TRANSACTION"

	LogErrMinioConn = "ERROR OPEN CONNECTION MINIO"
	LogErrMinioPut  = "ERROR MINIO PUT OBJECT"
	LogErrMinioDel  = "ERROR MINIO REMOVE OBJECT"

	LogErrFileCannotClose = "ERROR CANNOT CLOSE FILE HEADER"
	LogErrFileCannotOpen  = "ERROR CANNOT OPEN FILE HEADER"

	LogInfoDBTxRollback = "TRANSACTION ROLLBACK"
	LogInfoDBTxCommit   = "TRANSACTION COMMIT"
)
