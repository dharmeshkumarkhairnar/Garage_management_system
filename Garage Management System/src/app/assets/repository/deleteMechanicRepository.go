package repository

import (
	"context"
	"garage_management_system/src/app/assets/constants"
	"garage_management_system/src/app/assets/models"
	genModels "garage_management_system/src/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeleteMechanicRepository interface {
	DeleteMechanic(ctx context.Context, bffDeleteMechaniceRequest models.BFFDeleteMechanicRequest) error
}

type deleteMechanicRepository struct {
	gDB *gorm.DB
}

func NewDeleteMechanicRepository(gDB *gorm.DB) *deleteMechanicRepository {
	return &deleteMechanicRepository{gDB: gDB}
}

func (repo *deleteMechanicRepository) DeleteMechanic(ctx context.Context, bffDeleteMechaniceRequest models.BFFDeleteMechanicRequest) error {

	logger := logrus.New()

	result := repo.gDB.WithContext(ctx).Table(constants.MechanicsTableName).Where(constants.AadharNumberCondition, bffDeleteMechaniceRequest.AddharNumber).Delete(&genModels.Mechanics{})
	if result.Error != nil {
		return result.Error
	}

	logger.Info(constants.MechanicDeletedInDB)

	return nil
}
