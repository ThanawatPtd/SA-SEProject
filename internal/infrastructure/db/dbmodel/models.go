// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbmodel

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Employee struct {
	ID        pgtype.UUID        `json:"id"`
	Salary    float64            `json:"salary"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

type Transaction struct {
	ID            pgtype.UUID        `json:"id"`
	UserID        pgtype.UUID        `json:"userId"`
	VehicleID     pgtype.UUID        `json:"vehicleId"`
	Price         float64            `json:"price"`
	InsuranceType string             `json:"insuranceType"`
	Status        string             `json:"status"`
	ESlipImageUrl string             `json:"eSlipImageUrl"`
	CrImageUrl    string             `json:"crImageUrl"`
	CipNumber     pgtype.Text        `json:"cipNumber"`
	VipNumber     pgtype.Text        `json:"vipNumber"`
	CreatedAt     pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt     pgtype.Timestamptz `json:"updatedAt"`
}

type User struct {
	ID          pgtype.UUID        `json:"id"`
	Email       string             `json:"email"`
	Fname       string             `json:"fname"`
	Lname       string             `json:"lname"`
	Password    string             `json:"password"`
	PhoneNumber string             `json:"phoneNumber"`
	Address     string             `json:"address"`
	Nationality string             `json:"nationality"`
	BirthDate   pgtype.Timestamptz `json:"birthDate"`
	CitizenID   string             `json:"citizenId"`
	CreatedAt   pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt   pgtype.Timestamptz `json:"updatedAt"`
}

type Vehicle struct {
	ID                    pgtype.UUID        `json:"id"`
	RegistrationDate      pgtype.Timestamptz `json:"registrationDate"`
	RegistrationNumber    string             `json:"registrationNumber"`
	Province              string             `json:"province"`
	VehicleType           string             `json:"vehicleType"`
	VehicleCategory       string             `json:"vehicleCategory"`
	Characteristics       string             `json:"characteristics"`
	Brand                 string             `json:"brand"`
	Model                 string             `json:"model"`
	ModelYear             string             `json:"modelYear"`
	VehicleColor          string             `json:"vehicleColor"`
	VehicleNumber         string             `json:"vehicleNumber"`
	VehicleNumberLocation string             `json:"vehicleNumberLocation"`
	EngineBrand           string             `json:"engineBrand"`
	EngineNumber          string             `json:"engineNumber"`
	EngineNumberLocation  string             `json:"engineNumberLocation"`
	ChasisNumber          string             `json:"chasisNumber"`
	FuelType              string             `json:"fuelType"`
	WheelType             string             `json:"wheelType"`
	TotalPiston           int32              `json:"totalPiston"`
	Cc                    int32              `json:"cc"`
	HorsePower            int32              `json:"horsePower"`
	SeatingCapacity       int32              `json:"seatingCapacity"`
	WeightUnlanden        float64            `json:"weightUnlanden"`
	WeightLaden           float64            `json:"weightLaden"`
	Miles                 float64            `json:"miles"`
	CreatedAt             pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt             pgtype.Timestamptz `json:"updatedAt"`
}

type VehicleOwner struct {
	UserID    pgtype.UUID        `json:"userId"`
	VehicleID pgtype.UUID        `json:"vehicleId"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}
