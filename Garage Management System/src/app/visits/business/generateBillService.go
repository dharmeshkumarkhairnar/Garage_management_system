package business

import (
	"context"
	"errors"
	"fmt"
	"garage_management_system/src/app/visits/models"
	"garage_management_system/src/app/visits/repository"
)

type GenerateBillService struct {
	generateBillRepository repository.GenerateBillRepository
}

func NewGenerateBillervice(generateBillRepository repository.GenerateBillRepository) *GenerateBillService {
	return &GenerateBillService{
		generateBillRepository: generateBillRepository,
	}
}

func (service *GenerateBillService) GenerateBill(ctx context.Context, spanCtx context.Context, bffGenerateBillRequest models.BFFGenerateBillRequest) (float64, error) {

	serviceIds, err := service.generateBillRepository.GetServiceMasterID(ctx, bffGenerateBillRequest)
	if err != nil {
		return 0.0, fmt.Errorf("%w", err)
	}

	var data []uint64

	for i := 0; i < len(*serviceIds); i++ {
		data = append(data, (*serviceIds)[i].ServiceMasterID)
	}

	if len(data) == 0 {
		return 0.0, errors.New("Services not found for this ID")
	}

	billAmount, err := service.generateBillRepository.SumAmount(ctx, bffGenerateBillRequest, data)
	if err != nil {
		return 0.0, fmt.Errorf("%w", err)
	}

	return billAmount, nil
}
