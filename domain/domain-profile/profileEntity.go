package domainprofile

import "database/sql"

type Profile struct {
	ProfileId string
	UserId    string
	Quote     sql.NullString
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}
