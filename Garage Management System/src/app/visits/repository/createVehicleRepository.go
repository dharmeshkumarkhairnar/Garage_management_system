package repository

import (
	"context"
	"errors"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/models"
	genModels "garage_management_system/src/models"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateVehicleRepository interface {
	CreateVehicle(ctx context.Context, userID uint64, bffCreateVehicleRequest models.BFFCreateVehicleRequest) error
}

type createVehicleRepository struct {
	gDB *gorm.DB
}

func NewCreateVehicleRepository(gDB *gorm.DB) *createVehicleRepository {
	return &createVehicleRepository{gDB: gDB}
}

func (user *createVehicleRepository) CreateVehicle(ctx context.Context, userID uint64, bffCreateVehicleRequest models.BFFCreateVehicleRequest) error {

	logger := logrus.New()

	NewVehicle := genModels.Vehicles{
		UserID:  userID,
		NumberPlate: bffCreateVehicleRequest.NumberPlate,
		Model:       strings.ToLower(bffCreateVehicleRequest.Model),
		Created_at:  time.Now(),
	}

	result := user.gDB.WithContext(ctx).Table(constants.VehiclesTableName).Create(&NewVehicle)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), constants.DuplicateNumberPlateError) {
			return errors.New(constants.VehicleNumberPlateAlreadyExistsError)
		} else if strings.Contains(result.Error.Error(), constants.CustomerNotFoundError) {
			return errors.New(constants.UserNotFoundError)
		}
		return result.Error
	}

	logger.Info(constants.VehicleCreationSuccess)

	return nil
}
