package business

import (
	"context"
	"errors"
	"fmt"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/models"
	"garage_management_system/src/app/visits/repository"
	"strings"
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

	serviceIds, err := service.addVisitRecordRepository.GetServiceIds(ctx, bffAddVisitRecordsRequest)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	data := []map[string]interface{}{}

	for i := 0; i < len(*serviceIds); i++ {
		items := map[string]interface{}{constants.VisitRecordID: vistRecords.ID, constants.ServiceMasterID: (*serviceIds)[i].ID}
		data = append(data, items)
	}

	if len(bffAddVisitRecordsRequest.Services) != len(*serviceIds) {
		serviceString := fmt.Sprint(serviceIds)
		var errorMsg string
		for _, v := range bffAddVisitRecordsRequest.Services {
			if !strings.Contains(serviceString, v) {
				errorMsg = errorMsg + fmt.Sprintf("%s, ", v)
			}
		}
		errorMsg =constants.SomeServicesNotAvailableError + errorMsg
		return errors.New(errorMsg)
	}

	err = service.addVisitRecordRepository.AddVisitServices(ctx, data)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil

}
