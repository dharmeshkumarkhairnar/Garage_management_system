package repository

import (
	"context"
	"errors"
	"fmt"
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

type CreateCustomer interface {
	CreateNewCustomer(ctx context.Context, db *gorm.DB, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error
}

type createCustomer struct {
	// DB *gorm.DB
	logger *logrus.Logger
}

func NewCreateCustomer(db *gorm.DB, logger *logrus.Logger) *createCustomer {
	return &createCustomer{
		// DB: db,
		logger: logger,
	}
}

func (repo *createCustomer) CreateNewCustomer(ctx context.Context, db *gorm.DB, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error {

	start := time.Now()
	hashPassword, err := utils.HashPassword(bffCreateCustomerRequest.Password)
	if err != nil {
		log.Info(constants.ErrFailedToEncrypt)
		return err
	}

	bffCreateCustomerRequest.Password = hashPassword

	NewCustomer := genericModels.Customers{
		Name:      bffCreateCustomerRequest.Name,
		Email:     bffCreateCustomerRequest.Email,
		Phone:     bffCreateCustomerRequest.PhoneNumber,
		Password:  bffCreateCustomerRequest.Password,
		CreatedAt: time.Now(),
	}

	result := db.WithContext(ctx).Create(&NewCustomer)
	if result.Error != nil {
		errorMsgs := result.Error.Error()
		if strings.Contains(errorMsgs, constants.ErrUniqueConstraintViolation) {
			duplicateKeys := []string{}
			if strings.Contains(errorMsgs, constants.IndexUsersPanCard) {
				duplicateKeys = append(duplicateKeys, constants.FieldPanCard)
			}
			if strings.Contains(errorMsgs, constants.IndexUsersEmail) {
				duplicateKeys = append(duplicateKeys, constants.FieldEmail)
			}

			if len(duplicateKeys) > 0 {
				return errors.New(strings.Join(duplicateKeys, ",") + constants.ErrDuplicateEntry)
			}

			return errors.New(constants.ErrUsernameExists)
		}
		fmt.Println("err in repo:", result.Error)
		return result.Error
	}

	repo.logger.WithFields(logrus.Fields{
		"user":    bffCreateCustomerRequest.Email,
		"latency": time.Since(start).Milliseconds(),
	}).Info(constants.UserCreationSuccessMsg)

	return nil
}
