package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/assets/repository"
	"garage_management_system/src/app/assets/models"
)

type DeleteMechanicService struct {
	deleteMechanicRepository repository.DeleteMechanicRepository
}

func NewDeleteMechanicService(deleteMechanicRepository repository.DeleteMechanicRepository) *DeleteMechanicService {
	return &DeleteMechanicService{
		deleteMechanicRepository: deleteMechanicRepository,
	}
}

func (service *DeleteMechanicService) DeleteMechanic(ctx context.Context, spanCtx context.Context, bffDeleteMechanicRequest models.BFFDeleteMechanicRequest) error {

	err := service.deleteMechanicRepository.DeleteMechanic(spanCtx, bffDeleteMechanicRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
