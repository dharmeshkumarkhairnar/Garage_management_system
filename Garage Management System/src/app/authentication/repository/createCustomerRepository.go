package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"

	genericModels "garage_management_system/src/models"
	"garage_management_system/src/utils"

	"github.com/pingcap/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateCustomerRepository interface {
	CreateNewCustomer(ctx context.Context, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error
}

type createCustomerRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

func NewCreateCustomerRepository(db *gorm.DB, logger *logrus.Logger) *createCustomerRepository {
	return &createCustomerRepository{
		DB:     db,
		logger: logger,
	}
}

func (repo *createCustomerRepository) CreateNewCustomer(ctx context.Context, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error {

	start := time.Now()
	hashPassword, err := utils.HashPassword(bffCreateCustomerRequest.Password)
	if err != nil {
		log.Info(constants.ErrFailedToEncrypt)
		return err
	}

	bffCreateCustomerRequest.Password = hashPassword

	NewCustomer := genericModels.Users{
		Name:      bffCreateCustomerRequest.Name,
		Email:     bffCreateCustomerRequest.Email,
		Phone:     bffCreateCustomerRequest.PhoneNumber,
		Password:  bffCreateCustomerRequest.Password,
		CreatedAt: time.Now(),
	}

	result := repo.DB.WithContext(ctx).Create(&NewCustomer)
	if result.Error != nil {
		errorMsgs := result.Error.Error()

		if strings.Contains(errorMsgs, constants.ErrUniqueConstraintViolation) {
			duplicateKeys := []string{}

			if strings.Contains(errorMsgs, constants.IndexCustomersEmail) {
				duplicateKeys = append(duplicateKeys, constants.FieldEmail)
			}

			if len(duplicateKeys) > 0 {
				return errors.New(strings.Join(duplicateKeys, ",") + constants.ErrDuplicateEntry)
			}
		}

		return result.Error
	}

	repo.logger.WithFields(logrus.Fields{
		"user":    bffCreateCustomerRequest.Email,
		"latency": time.Since(start).Milliseconds(),
	}).Info(constants.UserCreationSuccessMsg)

	return nil
}
