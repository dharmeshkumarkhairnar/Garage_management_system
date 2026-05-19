package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/visits/models"
	"garage_management_system/src/app/visits/repository"
)

type CreateVehicleService struct {
	createVehicleRepository repository.CreateVehicleRepository
}

func NewCreateVehicleService(createVehicleRepository repository.CreateVehicleRepository) *CreateVehicleService {
	return &CreateVehicleService{
		createVehicleRepository: createVehicleRepository,
	}
}

func (service *CreateVehicleService) CreateVehicle(ctx context.Context, userID uint64, spanCtx context.Context, bffCreateVehicleRequest models.BFFCreateVehicleRequest) error {

	err := service.createVehicleRepository.CreateVehicle(spanCtx, userID, bffCreateVehicleRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
