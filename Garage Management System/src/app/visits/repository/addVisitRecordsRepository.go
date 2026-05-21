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
	AddVisitRecords(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) error
}

type addVisitRecordsRepository struct {
	gDB *gorm.DB
}

func NewAddVisitRecordsRepository(gDB *gorm.DB) *addVisitRecordsRepository {
	return &addVisitRecordsRepository{gDB: gDB}
}

func (user *addVisitRecordsRepository) AddVisitRecords(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) error {

	logger := logrus.New()

	var vehicles genericModels.Vehicles

	err := user.gDB.WithContext(ctx).Table(constants.VehiclesTableName).Where(constants.NumberPlate, bffAddVisitRecordsRequest.NumberPlate).First(&vehicles).Error
	if err != nil {
		return errors.New(constants.VehicleNotFoundError)
	}

	arrivalTime, _ := time.Parse("2026-05-20", bffAddVisitRecordsRequest.ArrivalDate)
	deliveryTime, _ := time.Parse("2026-05-20", bffAddVisitRecordsRequest.DeliveryDate)

	VisitRecord := genModels.VisitRecords{
		VehicleID:    vehicles.ID,
		MechanicID:   bffAddVisitRecordsRequest.MechanicId,
		ArrivalDate:  arrivalTime,
		DeliveryDate: deliveryTime,
	}

	result := user.gDB.WithContext(ctx).Clauses(clause.Returning{}).Table(constants.VisitRecordTableName).Create(&VisitRecord)
	if result.Error != nil {
		return errors.New("Database Insertion Error")
	}

	
	

	logger.Info("Visit Records Added Successfully")
	return nil
}

func (user *addVisitRecordsRepository) AddVisitServices(ctx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) ([]string, *genModels.VisitServices, error) {

	logger := logrus.New()

	var services genericModels.VisitServices

	err := user.gDB.WithContext(ctx).Find(&services).Error
	if err != nil {
		return nil, nil, errors.New("No Services Found")
	}

	logger.Info("Visit Records Added Successfully")
	return bffAddVisitRecordsRequest.Services, &services, nil
}
