package domainuser

import "database/sql"

// user entities
type User struct {
	ID              string
	FullName        string
	Gender          string
	Image           string
	Username        string
	Email           string
	Password        string
	PhoneNumber     sql.NullString
	EmailVerifiedAt bool
	CreatedAt       int64
	CreatedBy       string
	UpdatedAt       int64
	UpdatedBy       sql.NullString
	DeletedAt       sql.NullInt64
	DeletedBy       sql.NullString
}
