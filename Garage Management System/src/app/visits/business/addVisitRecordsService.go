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

	vistRecords, err := service.addVisitRecordRepository.AddVisitRecords(ctx, bffAddVisitRecordsRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println(vistRecords.ID)

	serviceIds, err := service.addVisitRecordRepository.GetServiceIds(ctx, bffAddVisitRecordsRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println(serviceIds)

	data := []map[string]interface{}{}

	for i := 0; i < len(*serviceIds); i++ {
		items := map[string]interface{}{"visit_record_id": vistRecords.ID, "service_master_id": (*serviceIds)[i].ID}
		data = append(data, items)
	}

	err = service.addVisitRecordRepository.AddVisitServices(ctx, data)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
