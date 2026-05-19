package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	"garage_management_system/src/app/authentication/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CreateCustomerService struct {
	createCustomerRepo repository.CreateCustomerRepository
	DB                 *gorm.DB
	logger             *logrus.Logger
}

func NewCreateUserService(createCustomerRepo repository.CreateCustomerRepository, db *gorm.DB, logger *logrus.Logger) *CreateCustomerService {
	return &CreateCustomerService{
		createCustomerRepo: createCustomerRepo,
		DB:                 db,
		logger:             logger,
	}
}

func (service *CreateCustomerService) CreateNewCustomer(spanCtx context.Context, bffCreateCustomerRequest models.BFFCreateCustomerRequest) error {
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

	return nil
}
