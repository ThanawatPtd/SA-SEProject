package repositories

import (
	"context"

	"github.com/ThanawatPtd/SAProject/domain/entities"
)

type VehicleRepository interface {
	CreateVehicle(ctx context.Context, vehicle *entities.Vehicle) (*entities.Vehicle, error)
	FindTemplate(ctx context.Context, userID string) ([]entities.Vehicle, error)
}
