// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: vehicle.sql

package dbmodel

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createVehicle = `-- name: CreateVehicle :one
INSERT INTO "vehicle" (
    registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year,vehicle_color, engine_number, chasis_number , fuel_type, horse_power, seating_capacity, weight_unlanden, weight_laden, tire_count, compulsory_insurance_policy_number, voluntary_insurance_policy_number, insurance_type, created_at, updated_at
) VALUES (
   $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, NOW(), NOW()
)
RETURNING vehicle_id,registration_date, registration_number, province, vehicle_type, vehicle_category, characteristics, brand, model, model_year,vehicle_color, engine_number, chasis_number , fuel_type, horse_power, seating_capacity, weight_unlanden, weight_laden, tire_count, compulsory_insurance_policy_number, voluntary_insurance_policy_number, insurance_type, created_at
`

type CreateVehicleParams struct {
	RegistrationDate                pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber              pgtype.Timestamptz `json:"registrationNumber"`
	Province                        string             `json:"province"`
	VehicleType                     string             `json:"vehicleType"`
	VehicleCategory                 string             `json:"vehicleCategory"`
	Characteristics                 string             `json:"characteristics"`
	Brand                           string             `json:"brand"`
	Model                           string             `json:"model"`
	ModelYear                       string             `json:"modelYear"`
	VehicleColor                    string             `json:"vehicleColor"`
	EngineNumber                    string             `json:"engineNumber"`
	ChasisNumber                    string             `json:"chasisNumber"`
	FuelType                        string             `json:"fuelType"`
	HorsePower                      int32              `json:"horsePower"`
	SeatingCapacity                 int32              `json:"seatingCapacity"`
	WeightUnlanden                  float64            `json:"weightUnlanden"`
	WeightLaden                     float64            `json:"weightLaden"`
	TireCount                       int32              `json:"tireCount"`
	CompulsoryInsurancePolicyNumber string             `json:"compulsoryInsurancePolicyNumber"`
	VoluntaryInsurancePolicyNumber  pgtype.Text        `json:"voluntaryInsurancePolicyNumber"`
	InsuranceType                   pgtype.Text        `json:"insuranceType"`
}

type CreateVehicleRow struct {
	VehicleID                       pgtype.UUID        `json:"vehicleId"`
	RegistrationDate                pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber              pgtype.Timestamptz `json:"registrationNumber"`
	Province                        string             `json:"province"`
	VehicleType                     string             `json:"vehicleType"`
	VehicleCategory                 string             `json:"vehicleCategory"`
	Characteristics                 string             `json:"characteristics"`
	Brand                           string             `json:"brand"`
	Model                           string             `json:"model"`
	ModelYear                       string             `json:"modelYear"`
	VehicleColor                    string             `json:"vehicleColor"`
	EngineNumber                    string             `json:"engineNumber"`
	ChasisNumber                    string             `json:"chasisNumber"`
	FuelType                        string             `json:"fuelType"`
	HorsePower                      int32              `json:"horsePower"`
	SeatingCapacity                 int32              `json:"seatingCapacity"`
	WeightUnlanden                  float64            `json:"weightUnlanden"`
	WeightLaden                     float64            `json:"weightLaden"`
	TireCount                       int32              `json:"tireCount"`
	CompulsoryInsurancePolicyNumber string             `json:"compulsoryInsurancePolicyNumber"`
	VoluntaryInsurancePolicyNumber  pgtype.Text        `json:"voluntaryInsurancePolicyNumber"`
	InsuranceType                   pgtype.Text        `json:"insuranceType"`
	CreatedAt                       pgtype.Timestamptz `json:"createdAt"`
}

func (q *Queries) CreateVehicle(ctx context.Context, arg CreateVehicleParams) (CreateVehicleRow, error) {
	row := q.db.QueryRow(ctx, createVehicle,
		arg.RegistrationDate,
		arg.RegistrationNumber,
		arg.Province,
		arg.VehicleType,
		arg.VehicleCategory,
		arg.Characteristics,
		arg.Brand,
		arg.Model,
		arg.ModelYear,
		arg.VehicleColor,
		arg.EngineNumber,
		arg.ChasisNumber,
		arg.FuelType,
		arg.HorsePower,
		arg.SeatingCapacity,
		arg.WeightUnlanden,
		arg.WeightLaden,
		arg.TireCount,
		arg.CompulsoryInsurancePolicyNumber,
		arg.VoluntaryInsurancePolicyNumber,
		arg.InsuranceType,
	)
	var i CreateVehicleRow
	err := row.Scan(
		&i.VehicleID,
		&i.RegistrationDate,
		&i.RegistrationNumber,
		&i.Province,
		&i.VehicleType,
		&i.VehicleCategory,
		&i.Characteristics,
		&i.Brand,
		&i.Model,
		&i.ModelYear,
		&i.VehicleColor,
		&i.EngineNumber,
		&i.ChasisNumber,
		&i.FuelType,
		&i.HorsePower,
		&i.SeatingCapacity,
		&i.WeightUnlanden,
		&i.WeightLaden,
		&i.TireCount,
		&i.CompulsoryInsurancePolicyNumber,
		&i.VoluntaryInsurancePolicyNumber,
		&i.InsuranceType,
		&i.CreatedAt,
	)
	return i, err
}

const getAllVehicle = `-- name: GetAllVehicle :many
SELECT
    vehicle_id,
    registration_date,
    registration_number,
    province,
    vehicle_type ,
    vehicle_category,
    characteristics ,
    brand,
    model ,
    model_year ,
    vehicle_color ,
    engine_number ,
    chasis_number ,
    fuel_type ,
    horse_power,
    seating_capacity,
    weight_unlanden ,
    weight_laden ,
    tire_count,
    compulsory_insurance_policy_number ,
    voluntary_insurance_policy_number ,
    insurance_type ,
    created_at,
    updated_at
from "vehicle"
`

func (q *Queries) GetAllVehicle(ctx context.Context) ([]Vehicle, error) {
	rows, err := q.db.Query(ctx, getAllVehicle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Vehicle
	for rows.Next() {
		var i Vehicle
		if err := rows.Scan(
			&i.VehicleID,
			&i.RegistrationDate,
			&i.RegistrationNumber,
			&i.Province,
			&i.VehicleType,
			&i.VehicleCategory,
			&i.Characteristics,
			&i.Brand,
			&i.Model,
			&i.ModelYear,
			&i.VehicleColor,
			&i.EngineNumber,
			&i.ChasisNumber,
			&i.FuelType,
			&i.HorsePower,
			&i.SeatingCapacity,
			&i.WeightUnlanden,
			&i.WeightLaden,
			&i.TireCount,
			&i.CompulsoryInsurancePolicyNumber,
			&i.VoluntaryInsurancePolicyNumber,
			&i.InsuranceType,
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
