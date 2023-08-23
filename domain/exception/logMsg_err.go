package exception

var (
	LogErrScanning    = "ERROR SCANNING DATA"
	LogErrSTMT        = "ERROR START PREPARED STATEMENT"
	LogErrExec        = "ERROR EXEC TO DATA"
	LogErrQuery       = "ERROR QUERY TO DATA"
	LogErrMinioConn   = "ERROR OPEN CONNECTION MINIO"
	LogErrMinioPut    = "ERROR MINIO PUT OBJECT"
	LogErrMinioDel    = "ERROR MINIO REMOVE OBJECT"
	LogErrStartTx     = "ERROR START TRANSACTION"
	LogErrTx          = "ERROR TRANSACTION"
	LogErrTxRollback  = "ERROR ROLLBACK TRANSACTION"
	LogErrTxCommit    = "ERROR COMMIT TRANSACTION"
	LogErrTxStart     = "CANNOT START TRANSACTION"
	LogInfoTxRollback = "TRANSACTION ROLLBACK"
	LogInfoTxCommit   = "TRANSACTION COMMIT"
)
