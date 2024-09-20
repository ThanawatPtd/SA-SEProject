// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbmodel

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Employee struct {
	ID          pgtype.UUID        `json:"id"`
	Email       string             `json:"email"`
	Fname       string             `json:"fname"`
	Lname       string             `json:"lname"`
	Password    string             `json:"password"`
	PhoneNumber string             `json:"phoneNumber"`
	Address     string             `json:"address"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
	Salary      float32            `json:"salary"`
	Position    string             `json:"position"`
}

type User struct {
	ID          pgtype.UUID        `json:"id"`
	Email       string             `json:"email"`
	Fname       string             `json:"fname"`
	Lname       string             `json:"lname"`
	Password    string             `json:"password"`
	PhoneNumber string             `json:"phoneNumber"`
	Address     string             `json:"address"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
}
