// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: insurance.sql

package dbmodel

import (
	"context"
)

const addInsuranceHeadcount = `-- name: AddInsuranceHeadcount :exec
UPDATE "insurance" 
SET headcount = headcount + 1
WHERE brand = $1 AND model = $2 AND year = $3
`

type AddInsuranceHeadcountParams struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
}

func (q *Queries) AddInsuranceHeadcount(ctx context.Context, arg AddInsuranceHeadcountParams) error {
	_, err := q.db.Exec(ctx, addInsuranceHeadcount, arg.Brand, arg.Model, arg.Year)
	return err
}

const createInsurance = `-- name: CreateInsurance :one
INSERT INTO "insurance" (
    year, model, brand, price
) VALUES (
    $1, $2, $3, $4
)
RETURNING year, model, brand, price
`

type CreateInsuranceParams struct {
	Year  string  `json:"year"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

type CreateInsuranceRow struct {
	Year  string  `json:"year"`
	Model string  `json:"model"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
}

func (q *Queries) CreateInsurance(ctx context.Context, arg CreateInsuranceParams) (CreateInsuranceRow, error) {
	row := q.db.QueryRow(ctx, createInsurance,
		arg.Year,
		arg.Model,
		arg.Brand,
		arg.Price,
	)
	var i CreateInsuranceRow
	err := row.Scan(
		&i.Year,
		&i.Model,
		&i.Brand,
		&i.Price,
	)
	return i, err
}

const deleteInsurance = `-- name: DeleteInsurance :exec
DELETE FROM "insurance"
WHERE brand = $1 AND model = $2 AND year = $3
`

type DeleteInsuranceParams struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
}

func (q *Queries) DeleteInsurance(ctx context.Context, arg DeleteInsuranceParams) error {
	_, err := q.db.Exec(ctx, deleteInsurance, arg.Brand, arg.Model, arg.Year)
	return err
}

const getInsurance = `-- name: GetInsurance :one
SELECT
    price
FROM "insurance"
WHERE brand = $1 AND model = $2 AND year = $3
`

type GetInsuranceParams struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  string `json:"year"`
}

func (q *Queries) GetInsurance(ctx context.Context, arg GetInsuranceParams) (float64, error) {
	row := q.db.QueryRow(ctx, getInsurance, arg.Brand, arg.Model, arg.Year)
	var price float64
	err := row.Scan(&price)
	return price, err
}

const getInsurances = `-- name: GetInsurances :many
SELECT
    model,
    brand,
    year
FROM "insurance"
ORDER BY brand, model, year
`

type GetInsurancesRow struct {
	Model string `json:"model"`
	Brand string `json:"brand"`
	Year  string `json:"year"`
}

func (q *Queries) GetInsurances(ctx context.Context) ([]GetInsurancesRow, error) {
	rows, err := q.db.Query(ctx, getInsurances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInsurancesRow
	for rows.Next() {
		var i GetInsurancesRow
		if err := rows.Scan(&i.Model, &i.Brand, &i.Year); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
