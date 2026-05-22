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
	var mechanic genModels.Mechanics

	err := repo.gDB.WithContext(ctx).Table(constants.VehiclesTableName).Where(constants.NumberPlate, bffAddVisitRecordsRequest.NumberPlate).First(&vehicles).Error
	if err != nil {
		logger.Error("error in getting the vehicle record")
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New(constants.VehicleNotFoundError)
		}
		return nil, err
	}
	logger.Info("Vehicle ID fetched successfully")

	err = repo.gDB.WithContext(ctx).Table(constants.MechanicsTableName).Where(constants.MechIDCondition, bffAddVisitRecordsRequest.MechanicId).First(&mechanic).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("error in getting the mechanic record")
			return nil, errors.New(constants.MechanicNotFoundError)
		}
		return nil, err
	}
	logger.Info("mechanic record fetched successfully")

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
		logger.Error("error in creating the visit record")
		return nil, result.Error
	}

	logger.Info("Visit Records Added Successfully")
	return &VisitRecord, nil
}

func (repo *addVisitRecordsRepository) GetServiceIds(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) (*[]genModels.ServiceMaster, error) {

	logger := logrus.New()

	var services []genericModels.ServiceMaster

	err := repo.gDB.WithContext(ctx).
		Table(constants.ServiceMasterTableName).
		Where(constants.ServiceIDINCondition, bffAddVisitRecordsRequest.Services).
		Find(&services).Error

	if err != nil {
		logger.Error("error in fetching service records")
		return nil, err
	}

	if len(services) == 0 {
		logger.Error("empty slice: No records found")
		return nil, errors.New(constants.ServiceNotFoundError)
	}

	logger.Info("Visit Services Fetched Successfully")
	return &services, nil
}

func (repo *addVisitRecordsRepository) AddVisitServices(ctx context.Context, data []map[string]interface{}) error {

	logger := logrus.New()

	err := repo.gDB.WithContext(ctx).
		Table(constants.VisitServiceTableName).Create(data).Error

	if err != nil {
		logger.Error("error in creating the visit service record")
		return err
	}

	logger.Info("Visit services addded Successfully")
	return nil
}
