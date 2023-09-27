// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"database/sql"
)

type User struct {
	ID               int64          `json:"id"`
	Email            string         `json:"email"`
	PasswordHash     string         `json:"password_hash"`
	Verified         bool           `json:"verified"`
	VerificationCode sql.NullString `json:"verification_code"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	UpdatedAt        sql.NullTime   `json:"updated_at"`
}
