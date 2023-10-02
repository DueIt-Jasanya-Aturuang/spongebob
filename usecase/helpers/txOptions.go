package helpers

import (
	"database/sql"
)

func LevelReadCommitted() *sql.TxOptions {
	return &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
		ReadOnly:  false,
	}
}
