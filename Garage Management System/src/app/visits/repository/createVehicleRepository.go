package repository

import (
	"context"
	"garage_management_system/src/app/visits/constants"
	"garage_management_system/src/app/visits/models"
	genModels "garage_management_system/src/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateVehicleRepository interface {
	CreateVehicle(ctx context.Context, bffCreateVehicleRequest models.BFFCreateVehicleRequest) error
}

type createVehicleRepository struct {
	gDB *gorm.DB
}

func NewCreateVehicleRepository(gDB *gorm.DB) *createVehicleRepository {
	return &createVehicleRepository{gDB: gDB}
}

func (user *createVehicleRepository) CreateVehicle(ctx context.Context, bffCreateVehicleRequest models.BFFCreateVehicleRequest) error {

	logger := logrus.New()

	NewVehicle := genModels.Vehicles{
		CustomerID:  bffCreateVehicleRequest.CustomerID,
		NumberPlate: bffCreateVehicleRequest.NumberPlate,
		Model:       bffCreateVehicleRequest.Model,
		Created_at:  time.Now(),
	}

	result := user.gDB.WithContext(ctx).Table(constants.TableName).Create(&NewVehicle)
	if result.Error != nil {
		return result.Error
	}

	logger.Info(constants.VehicleCreationSuccess)

	return nil
}
