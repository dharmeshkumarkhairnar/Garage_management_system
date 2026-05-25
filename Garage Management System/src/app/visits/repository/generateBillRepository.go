package repository

import (
	"context"
	"errors"
	"fmt"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/models"
	genModels "garage_management_system/src/models"
	genericModels "garage_management_system/src/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GenerateBillRepository interface {
	GetServiceMasterID(ctx context.Context, bffGenerateBillRequest models.BFFGenerateBillRequest) (*[]genModels.VisitServices, error)
	SumAmount(ctx context.Context, bffGenerateBillRequest models.BFFGenerateBillRequest, serviceIds []uint64) (float64, error)
}

type generateBillRepository struct {
	gDB *gorm.DB
}

func NewGenerateBillRepository(gDB *gorm.DB) *generateBillRepository {
	return &generateBillRepository{gDB: gDB}
}

func (repo *generateBillRepository) GetServiceMasterID(ctx context.Context, bffGenerateBillRequest models.BFFGenerateBillRequest) (*[]genModels.VisitServices, error) {

	logger := logrus.New()

	var services []genericModels.VisitServices

	err := repo.gDB.WithContext(ctx).
		Table(constants.VisitServiceTableName).
		Where(constants.FieldVisitRecordID, bffGenerateBillRequest.VisitRecordID). // Matches all names in the list
		Find(&services).Error

	fmt.Println(services)
	if err != nil {
		return nil, errors.New("Record Not Found")
	}

	logger.Info("Service master IDs Fetched Successfully")
	return &services, nil
}

func (repo *generateBillRepository) SumAmount(ctx context.Context, bffGenerateBillRequest models.BFFGenerateBillRequest, serviceIds []uint64) (float64, error) {

	logger := logrus.New()

	var amount float64

	err := repo.gDB.WithContext(ctx).
		Table(constants.ServiceMasterTableName).
		Select("SUM(amount) as total_amount").
		Where("id IN ?", serviceIds).
		Find(&amount).Error

	if err != nil {
		return 0.0, errors.New("Record Not Found")
	}

	logger.Info("Sum Amount Successfully")
	return amount, nil
}
