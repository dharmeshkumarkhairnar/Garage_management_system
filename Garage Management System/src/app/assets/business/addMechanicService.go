package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/assets/models"
	"garage_management_system/src/app/assets/repository"
)

type AddMechanicService struct {
	addMechanicRepository repository.AddMechanicRepository
}

func NewAddMechanicService(addMechanicRepository repository.AddMechanicRepository) *AddMechanicService {
	return &AddMechanicService{
		addMechanicRepository: addMechanicRepository,
	}
}

func (service *AddMechanicService) AddMechanic(ctx context.Context, spanCtx context.Context, bffAddMechanicRequest models.BFFAddMechanicRequest) error {

	err := service.addMechanicRepository.AddMechanic(spanCtx, bffAddMechanicRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
