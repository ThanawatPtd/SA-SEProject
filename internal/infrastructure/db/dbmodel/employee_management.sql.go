// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: employee_management.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEmployeeManagement = `-- name: CreateEmployeeManagement :one
INSERT INTO "employee_management" (
    employee_id, admin_id, created_at, updated_at
) VALUES (
    $1, $2, NOW(), NOW()
)
RETURNING employee_id, admin_id, created_at
`

type CreateEmployeeManagementParams struct {
	EmployeeID pgtype.UUID `json:"employeeId"`
	AdminID    pgtype.UUID `json:"adminId"`
}

type CreateEmployeeManagementRow struct {
	EmployeeID pgtype.UUID        `json:"employeeId"`
	AdminID    pgtype.UUID        `json:"adminId"`
	CreatedAt  pgtype.Timestamptz `json:"createdAt"`
}

func (q *Queries) CreateEmployeeManagement(ctx context.Context, arg CreateEmployeeManagementParams) (CreateEmployeeManagementRow, error) {
	row := q.db.QueryRow(ctx, createEmployeeManagement, arg.EmployeeID, arg.AdminID)
	var i CreateEmployeeManagementRow
	err := row.Scan(&i.EmployeeID, &i.AdminID, &i.CreatedAt)
	return i, err
}

const deleteEmployeeManagement = `-- name: DeleteEmployeeManagement :exec
DELETE FROM "employee_management"
WHERE employee_id = $1 AND admin_id = $2
`

type DeleteEmployeeManagementParams struct {
	EmployeeID pgtype.UUID `json:"employeeId"`
	AdminID    pgtype.UUID `json:"adminId"`
}

func (q *Queries) DeleteEmployeeManagement(ctx context.Context, arg DeleteEmployeeManagementParams) error {
	_, err := q.db.Exec(ctx, deleteEmployeeManagement, arg.EmployeeID, arg.AdminID)
	return err
}

const getAllEmployeeManagement = `-- name: GetAllEmployeeManagement :many
SELECT
    employee_id,
    admin_id,
    created_at,
    updated_at
FROM "employee_management"
`

func (q *Queries) GetAllEmployeeManagement(ctx context.Context) ([]EmployeeManagement, error) {
	rows, err := q.db.Query(ctx, getAllEmployeeManagement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmployeeManagement
	for rows.Next() {
		var i EmployeeManagement
		if err := rows.Scan(
			&i.EmployeeID,
			&i.AdminID,
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

const getAllEmployeeManagementsByAdminID = `-- name: GetAllEmployeeManagementsByAdminID :many
SELECT
   employee_id,
   admin_id,
   created_at,
   updated_at 
FROM "employee_management"
WHERE admin_id = $1
`

func (q *Queries) GetAllEmployeeManagementsByAdminID(ctx context.Context, adminID pgtype.UUID) ([]EmployeeManagement, error) {
	rows, err := q.db.Query(ctx, getAllEmployeeManagementsByAdminID, adminID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmployeeManagement
	for rows.Next() {
		var i EmployeeManagement
		if err := rows.Scan(
			&i.EmployeeID,
			&i.AdminID,
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

const getEmployeeManagementsByEmployeeID = `-- name: GetEmployeeManagementsByEmployeeID :many
SELECT
   employee_id,
   admin_id,
   created_at,
   updated_at 
FROM "employee_management"
WHERE employee_id = $1
`

func (q *Queries) GetEmployeeManagementsByEmployeeID(ctx context.Context, employeeID pgtype.UUID) ([]EmployeeManagement, error) {
	rows, err := q.db.Query(ctx, getEmployeeManagementsByEmployeeID, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EmployeeManagement
	for rows.Next() {
		var i EmployeeManagement
		if err := rows.Scan(
			&i.EmployeeID,
			&i.AdminID,
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

const updateEmployeeManagement = `-- name: UpdateEmployeeManagement :one
UPDATE "employee_management"
SET
    updated_at = NOW()
WHERE employee_id = $1 AND admin_id = $2
RETURNING employee_id, admin_id, updated_at
`

type UpdateEmployeeManagementParams struct {
	EmployeeID pgtype.UUID `json:"employeeId"`
	AdminID    pgtype.UUID `json:"adminId"`
}

type UpdateEmployeeManagementRow struct {
	EmployeeID pgtype.UUID        `json:"employeeId"`
	AdminID    pgtype.UUID        `json:"adminId"`
	UpdatedAt  pgtype.Timestamptz `json:"updatedAt"`
}

func (q *Queries) UpdateEmployeeManagement(ctx context.Context, arg UpdateEmployeeManagementParams) (UpdateEmployeeManagementRow, error) {
	row := q.db.QueryRow(ctx, updateEmployeeManagement, arg.EmployeeID, arg.AdminID)
	var i UpdateEmployeeManagementRow
	err := row.Scan(&i.EmployeeID, &i.AdminID, &i.UpdatedAt)
	return i, err
}