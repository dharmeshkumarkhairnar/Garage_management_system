package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	"garage_management_system/src/app/assets/repository"

	"gorm.io/gorm"
)

type AddNewServiceService struct {
	repo repository.AddNewServiceRepository
	db   *gorm.DB
}

func NewAddNewServiceService(repo repository.AddNewServiceRepository, db *gorm.DB) *AddNewServiceService {
	return &AddNewServiceService{
		repo: repo,
		db:   db,
	}
}

func (service *AddNewServiceService) AddNewService(ctx context.Context, bffAddNewServiceRequest models.BFFAddNewServiceRequest) error {
	// start := time.Now()

	tx := service.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf(constants.ErrBeginTx, tx.Error)
	}

	err := service.repo.AddNewService(ctx, tx, bffAddNewServiceRequest)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%w", err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf(constants.ErrCommitTx, err)
	}

	return nil
}
