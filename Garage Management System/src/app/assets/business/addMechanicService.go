package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/assets/models"
	"garage_management_system/src/app/assets/repository"
	"math/rand/v2"
	"strconv"
	"strings"
)

type AddMechanicService struct {
	addMechanicRepository repository.AddMechanicRepository
}

func NewAddMechanicService(addMechanicRepository repository.AddMechanicRepository) *AddMechanicService {
	return &AddMechanicService{
		addMechanicRepository: addMechanicRepository,
	}
}

func (service *AddMechanicService) AddMechanic(ctx context.Context, spanCtx context.Context, bffAddMechanicRequest models.BFFAddMechanicRequest) (string, error) {

	name := strings.Split(bffAddMechanicRequest.Name, " ")
	randomNum := rand.IntN(900) + 100

	mechanicID := strings.ToLower(name[0]) + strconv.Itoa(randomNum)

	err := service.addMechanicRepository.AddMechanic(spanCtx, mechanicID, bffAddMechanicRequest)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return mechanicID, nil

}
