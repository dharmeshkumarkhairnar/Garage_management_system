package repository

import (
	"context"
	"garage_management_system/src/app/assets/constants"
	"garage_management_system/src/app/assets/models"
	genModels "garage_management_system/src/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddMechanicRepository interface {
	AddMechanic(ctx context.Context, bffAddMechaniceRequest models.BFFAddMechanicRequest) error
}

type addMechanicRepository struct {
	gDB *gorm.DB
}

func NewAddMechanicRepository(gDB *gorm.DB) *addMechanicRepository {
	return &addMechanicRepository{gDB: gDB}
}

func (user *addMechanicRepository) AddMechanic(ctx context.Context, bffAddMechaniceRequest models.BFFAddMechanicRequest) error {

	logger := logrus.New()

	NewMechanic := genModels.Mechanics{
		Name:         bffAddMechaniceRequest.Name,
		AadharNumber: bffAddMechaniceRequest.AddharNumber,
		Phone:        bffAddMechaniceRequest.Phone,
		Created_at:   time.Now(),
	}

	result := user.gDB.WithContext(ctx).Table(constants.MechanicsTableName).Create(&NewMechanic)
	if result.Error != nil {
		// if strings.Contains(result.Error.Error(), constants.DuplicateNumberPlateError) {
		// 	return errors.New(constants.VehicleNumberPlateAlreadyExistsError)
		// } else if strings.Contains(result.Error.Error(), constants.CustomerNotFoundError) {
		// 	return errors.New(constants.UserNotFoundError)
		// }
		return result.Error
	}

	logger.Info(constants.MechanicAddedInDB)

	return nil
}
