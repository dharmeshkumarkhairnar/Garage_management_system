package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	"garage_management_system/src/app/authentication/repository"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateCustomer struct {
	createCustomerRepo repository.CreateCustomer
	DB                 *gorm.DB
	logger             *logrus.Logger
}

func NewCreateUserService(createCustomerRepo repository.CreateCustomer, db *gorm.DB, logger *logrus.Logger) *CreateCustomer {
	return &CreateCustomer{
		createCustomerRepo: createCustomerRepo,
		DB:                 db,
		logger:             logger,
	}
}

func (service *CreateCustomer) CreateNewCustomer(spanCtx context.Context, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error {
	start := time.Now()
	tx := service.DB.Begin()

	if tx.Error != nil {
		fmt.Println("transaction BEGIN error ", tx.Error)
		return fmt.Errorf(constants.ErrBeginTx, tx.Error)
	}

	err := service.createCustomerRepo.CreateNewCustomer(spanCtx, tx, bffCreateCustomerRequest)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%w", err)
	}

	if err := tx.Commit().Error; err != nil {
		fmt.Println("transaction COMMIT error ", err)
		return fmt.Errorf(constants.ErrCommitTx, err)
	}

	service.logger.WithFields(logrus.Fields{
		"user":    bffCreateCustomerRequest.Email,
		"latency": time.Since(start).Milliseconds(),
	}).Info("user creation transaction successful")

	return nil
}
