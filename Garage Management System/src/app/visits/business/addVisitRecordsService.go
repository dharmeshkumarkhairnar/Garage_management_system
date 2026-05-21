package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/visits/models"
	"garage_management_system/src/app/visits/repository"
)

type AddVisitRecordService struct {
	addVisitRecordRepository repository.AddVisitRecordsRepository
}

func NewAddVisitRecordService(addVisitRecordRepository repository.AddVisitRecordsRepository) *AddVisitRecordService {
	return &AddVisitRecordService{
		addVisitRecordRepository: addVisitRecordRepository,
	}
}

func (service *AddVisitRecordService) AddVisitRecord(ctx context.Context, spanCtx context.Context, bffAddVisitRecordsRequest models.BFFAddVisitRecordsRequest) error {

	err := service.addVisitRecordRepository.AddVisitRecords(ctx,bffAddVisitRecordsRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
