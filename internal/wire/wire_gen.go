// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/usecases"
	"github.com/ThanawatPtd/SAProject/internal/adapters/psql"
	"github.com/ThanawatPtd/SAProject/internal/adapters/rest"
	"github.com/ThanawatPtd/SAProject/internal/infrastructure/db"
)

import (
	_ "github.com/ThanawatPtd/SAProject/config"
)

// Injectors from wire.go:

func InitializeHandler() *rest.Handler {
	context := ProvideContext()
	configConfig := config.ProvideConfig()
	pool := db.ProvidePgxPool(context, configConfig)
	userRepository := psql.ProvidePostgresUserRepository(pool)
	userUseCase := usecases.ProvideUserService(userRepository, configConfig)
	userRestHandler := rest.ProvideUserRestHandler(userUseCase)
	employeeRepository := psql.ProvidePostgresEmployeeRepository(pool)
	employeeUseCase := usecases.ProvideEmployeeService(employeeRepository)
	employeeHandler := rest.ProvideEmployeeRestHandler(employeeUseCase)
	vehicleRepository := psql.ProvidePostgresVehicleRepository(pool)
	vehicleOwnerRepository := psql.ProvidePostgresVehicleOwnerRepository(pool)
	vehicleUseCase := usecases.ProvideVehicleService(vehicleRepository, vehicleOwnerRepository)
	vehicleRestHandler := rest.ProvideVehicleHandler(vehicleUseCase)
	transactionRepository := psql.ProvidePostgresTransactionRepository(pool)
	transactionUseCase := usecases.ProvideTransactionService(transactionRepository, vehicleOwnerRepository)
	transactionRestHandler := rest.ProvideTransactionRestHandler(transactionUseCase)
	authUseCase := usecases.ProvideAuthService(userRepository, employeeRepository, configConfig)
	authHandler := rest.ProvideAuthRestHandler(authUseCase)
	handler := rest.ProvideHandler(userRestHandler, employeeHandler, vehicleRestHandler, transactionRestHandler, authHandler)
	return handler
}
