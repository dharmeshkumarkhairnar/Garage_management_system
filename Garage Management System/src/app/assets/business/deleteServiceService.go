package business

import (
	"fmt"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/repository"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type DeleteServiceService struct {
	repo repository.DeleteServiceRepository
	db   *gorm.DB
}

func NewDeleteServiceService(repo repository.DeleteServiceRepository, db *gorm.DB) *DeleteServiceService {
	return &DeleteServiceService{
		repo: repo,
		db:   db,
	}
}

func (service *DeleteServiceService) DeleteService(ctx context.Context, serviceName string) error {

	tx := service.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf(constants.ErrBeginTx, tx.Error)
	}

	err := service.repo.DeleteService(ctx, serviceName)
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
