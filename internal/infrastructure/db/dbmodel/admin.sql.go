// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: admin.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAdmin = `-- name: CreateAdmin :one
INSERT INTO "admin" (
    id, created_at, updated_at
) VALUES (
    $1, NOW(), NOW()
)
RETURNING id, created_at
`

type CreateAdminRow struct {
	ID        pgtype.UUID        `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
}

func (q *Queries) CreateAdmin(ctx context.Context, id pgtype.UUID) (CreateAdminRow, error) {
	row := q.db.QueryRow(ctx, createAdmin, id)
	var i CreateAdminRow
	err := row.Scan(&i.ID, &i.CreatedAt)
	return i, err
}

const deleteAdmin = `-- name: DeleteAdmin :exec
DELETE FROM "admin"
WHERE id = $1
`

func (q *Queries) DeleteAdmin(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteAdmin, id)
	return err
}

const getAdminByID = `-- name: GetAdminByID :one
SELECT
    id,
    created_at,
    updated_at
FROM "admin"
WHERE id = $1
`

func (q *Queries) GetAdminByID(ctx context.Context, id pgtype.UUID) (Admin, error) {
	row := q.db.QueryRow(ctx, getAdminByID, id)
	var i Admin
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}

const getAllAdmins = `-- name: GetAllAdmins :many
SELECT
    id,
    created_at,
    updated_at
FROM "admin"
`

func (q *Queries) GetAllAdmins(ctx context.Context) ([]Admin, error) {
	rows, err := q.db.Query(ctx, getAllAdmins)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Admin
	for rows.Next() {
		var i Admin
		if err := rows.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdmin = `-- name: UpdateAdmin :one
UPDATE "admin"
SET
    updated_at = NOW()
WHERE id = $1
RETURNING id, updated_at
`

type UpdateAdminRow struct {
	ID        pgtype.UUID        `json:"id"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) UpdateAdmin(ctx context.Context, id pgtype.UUID) (UpdateAdminRow, error) {
	row := q.db.QueryRow(ctx, updateAdmin, id)
	var i UpdateAdminRow
	err := row.Scan(&i.ID, &i.UpdatedAt)
	return i, err
}