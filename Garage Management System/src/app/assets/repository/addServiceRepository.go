package repository

import (
	"errors"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genericModels "garage_management_system/src/models"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type addNewServiceRepository struct {
	// DB *gorm.DB
	logger *logrus.Logger
}

type AddNewServiceRepository interface {
	AddNewService(ctx context.Context, db *gorm.DB, bffAddNewServiceRequest models.BFFAddNewServiceRequest) error
}

func NewAddNewServiceRepository(logger *logrus.Logger) *addNewServiceRepository {
	return &addNewServiceRepository{
		logger: logger,
	}
}

func (repo *addNewServiceRepository) AddNewService(ctx context.Context, db *gorm.DB, bffAddNewServiceRequest models.BFFAddNewServiceRequest) error {
	start := time.Now()
	newService := genericModels.ServiceMaster{
		Service: strings.ToLower(bffAddNewServiceRequest.Service),
		Amount:  bffAddNewServiceRequest.Amount,
	}

	result := db.WithContext(ctx).Create(&newService)
	if result.Error != nil {
		errMsgs := result.Error.Error()

		if strings.Contains(errMsgs, constants.ErrUniqueConstraintViolation) {
			duplicateKeys := []string{}

			if strings.Contains(errMsgs, constants.IndexServiceName) {
				duplicateKeys = append(duplicateKeys, constants.IndexServiceName)
			}

			if len(duplicateKeys) > 0 {
				return errors.New(constants.ErrServiceAlreadyExists)
			}
		}

		return result.Error
	}

	repo.logger.WithFields(logrus.Fields{
		"service": bffAddNewServiceRequest.Service,
		"latency": time.Since(start).Milliseconds(),
	}).Info(constants.ServiceAddedSuccessfully)
	return nil
}
