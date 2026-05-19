package repository

import (
	"context"
	"errors"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genModels "garage_management_system/src/models"
	"strings"
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

func (repo *addMechanicRepository) AddMechanic(ctx context.Context, bffAddMechaniceRequest models.BFFAddMechanicRequest) error {

	logger := logrus.New()

	NewMechanic := genModels.Mechanics{
		Name:         bffAddMechaniceRequest.Name,
		AadharNumber: bffAddMechaniceRequest.AddharNumber,
		Phone:        bffAddMechaniceRequest.Phone,
		Created_at:   time.Now(),
	}

	result := repo.gDB.WithContext(ctx).Table(constants.MechanicsTableName).Create(&NewMechanic)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), constants.DuplicateAadharNumberDBError) {
			return errors.New(constants.DuplicateAadharNumberError)
		}
		return result.Error
	}

	logger.Info(constants.MechanicAddedInDB)

	return nil
}
