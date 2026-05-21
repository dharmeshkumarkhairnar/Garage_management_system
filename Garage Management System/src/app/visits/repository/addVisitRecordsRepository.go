package repository

import (
	"context"
	"errors"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/models"
	genModels "garage_management_system/src/models"
	genericModels "garage_management_system/src/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AddVisitRecordsRepository interface {
	AddVisitRecords(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) (*genModels.VisitRecords, error)
	GetServiceIds(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) (*[]genModels.ServiceMaster, error)
	AddVisitServices(ctx context.Context, data []map[string]interface{}) error
}

type addVisitRecordsRepository struct {
	gDB *gorm.DB
}

func NewAddVisitRecordsRepository(gDB *gorm.DB) *addVisitRecordsRepository {
	return &addVisitRecordsRepository{gDB: gDB}
}

func (repo *addVisitRecordsRepository) AddVisitRecords(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) (*genModels.VisitRecords, error) {

	logger := logrus.New()

	var vehicles genericModels.Vehicles

	err := repo.gDB.WithContext(ctx).Table(constants.VehiclesTableName).Where(constants.NumberPlate, bffAddVisitRecordsRequest.NumberPlate).First(&vehicles).Error
	if err != nil {
		return nil, errors.New(constants.VehicleNotFoundError)
	}

	arrivalTime, _ := time.Parse(time.DateOnly, bffAddVisitRecordsRequest.ArrivalDate)
	deliveryTime, _ := time.Parse(time.DateOnly, bffAddVisitRecordsRequest.DeliveryDate)

	VisitRecord := genModels.VisitRecords{
		VehicleID:    vehicles.ID,
		MechanicID:   bffAddVisitRecordsRequest.MechanicId,
		ArrivalDate:  arrivalTime,
		DeliveryDate: deliveryTime,
	}

	result := repo.gDB.WithContext(ctx).Clauses(clause.Returning{}).Table(constants.VisitRecordTableName).Create(&VisitRecord)
	if result.Error != nil {
		return nil, errors.New("Database Insertion Error")
	}

	logger.Info("Visit Records Added Successfully")
	return &VisitRecord, nil
}

func (repo *addVisitRecordsRepository) GetServiceIds(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) (*[]genModels.ServiceMaster, error) {

	logger := logrus.New()

	var services []genericModels.ServiceMaster

	err := repo.gDB.WithContext(ctx).
		Table(constants.ServiceMasterTableName).
		Where("service IN ?", bffAddVisitRecordsRequest.Services). // Matches all names in the list
		Find(&services).Error

	if err != nil {
		return nil, errors.New("Record Not Found")
	}

	logger.Info("Visit ServiceId Fetched Successfully")
	return &services, nil
}

func (repo *addVisitRecordsRepository) AddVisitServices(ctx context.Context, data []map[string]interface{}) error {

	logger := logrus.New()

	err := repo.gDB.WithContext(ctx).
		Table(constants.VisitServiceTableName).Create(data).Error

	if err != nil {
		return errors.New("Record Not Found")
	}

	logger.Info("Visit services addded Successfully")
	return nil
}
