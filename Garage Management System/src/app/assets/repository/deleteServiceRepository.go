package repository

import (
	"context"
	"fmt"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type deleteServiceRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

type DeleteServiceRepository interface {
	DeleteService(ctx context.Context, serviceName string) error
}

func NewDeleteServiceRepository(logger *logrus.Logger, db *gorm.DB) *deleteServiceRepository {
	return &deleteServiceRepository{
		DB:     db,
		logger: logger,
	}
}

func (repo *deleteServiceRepository) DeleteService(ctx context.Context, serviceName string) error {
	start := time.Now()
	result := repo.DB.WithContext(ctx).Where("service = ?", serviceName).Delete(&models.ServiceMaster{})
	if result.Error != nil {
		fmt.Println("Repo ERROR: ", result.Error)
		errMsgs := result.Error
		// if strings.Contains(errMsgs.Error(), gorm.ErrRecordNotFound.Error()) {
		// 	return errors.New("service does not exist")
		// }
		return errMsgs
	}

	repo.logger.WithFields(logrus.Fields{
		"service": serviceName,
		"latency": time.Since(start),
	}).Info(constants.ServiceDeletedSuccessfully)
	return nil
}
