// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" (
    id, email, fname, lname, password, phone_number, address, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
)
RETURNING id, fname AS name, email, created_at, updated_at
`

type CreateUserParams struct {
	ID          pgtype.UUID `json:"id"`
	Email       string      `json:"email"`
	Fname       string      `json:"fname"`
	Lname       string      `json:"lname"`
	Password    string      `json:"password"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     string      `json:"address"`
}

type CreateUserRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.Fname,
		arg.Lname,
		arg.Password,
		arg.PhoneNumber,
		arg.Address,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT
    id,
    fname AS name,
    email,
    created_at,
    updated_at
FROM "user"
`

type GetAllUsersRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByID = `-- name: GetUserByID :one
SELECT
    id,
    fname AS name,
    email,
    created_at,
    updated_at
FROM "user"
WHERE id = $1
`

type GetUserByIDRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) GetUserByID(ctx context.Context, id pgtype.UUID) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE "user"
SET
    email = $2,
    fname = $3,
    lname = $4,
    password = $5,
    phone_number = $6,
    address = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING id, fname AS name, email, created_at, updated_at
`

type UpdateUserParams struct {
	ID          pgtype.UUID `json:"id"`
	Email       string      `json:"email"`
	Fname       string      `json:"fname"`
	Lname       string      `json:"lname"`
	Password    string      `json:"password"`
	PhoneNumber string      `json:"phoneNumber"`
	Address     string      `json:"address"`
}

type UpdateUserRow struct {
	ID        pgtype.UUID        `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Fname,
		arg.Lname,
		arg.Password,
		arg.PhoneNumber,
		arg.Address,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
