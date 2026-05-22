package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/commons/constants"
	visitModels "garage_management_system/src/app/visits/models"
	"garage_management_system/src/models"
	commonModels "garage_management_system/src/models"
	"garage_management_system/src/utils/validations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddVisitRecordHandler struct {
	service *business.AddVisitRecordService
}

func NewAddVisitRecordHandler(service *business.AddVisitRecordService) *AddVisitRecordHandler {
	return &AddVisitRecordHandler{
		service: service,
	}
}

// HandlerAddVisitRecord handles the visit addition request.
// @Summary Create a new visit record
// @Description Handles visit addition request by validating input and storing visit details
// @Tags Visits
// @Accept json
// @Produce json
// @Param request body models.BFFAddVisitRecordsRequest true "Vehicle Registration Request"
// @Success 201 {object} models.BFFAddVisitRecordsResponse "User created successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 404 {object} models.ErrorAPIResponse "User not found"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/visits/add-record [post]
func (controller *AddVisitRecordHandler) AddVisitRecord(ctx *gin.Context) {

	var bffAddVisitRecordsRequest visitModels.BFFAddVisitRecordsRequest

	if err := ctx.ShouldBind(&bffAddVisitRecordsRequest); err != nil {
		errorMsgs := commonModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, commonModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	if err := validations.GetBFFValidator().Struct(&bffAddVisitRecordsRequest); err != nil {
		validationErros, _ := validations.FormatValidationErrors(err)
		ctx.IndentedJSON(http.StatusBadRequest, validationErros)
		return
	}

	err := controller.service.AddVisitRecord(ctx, ctx.Request.Context(), bffAddVisitRecordsRequest)
	if err != nil {
		if strings.Contains(err.Error(), constants.VehicleNotFoundError) {
			errorMsg := models.ErrorMessage{Key: constants.Vehicle, ErrorMessage: constants.VehicleNotFoundError}
			errorResponse := commonModels.ErrorAPIResponse{
				Message: errorMsg,
				Error:   constants.VisitRecordCreationFailedError,
			}
			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		} else if strings.Contains(err.Error(), constants.MechanicNotFoundError) {
			errorMsg := models.ErrorMessage{Key: constants.Mechanic, ErrorMessage: constants.MechanicNotFoundError}
			errorResponse := commonModels.ErrorAPIResponse{
				Message: errorMsg,
				Error:   constants.VisitRecordCreationFailedError,
			}
			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		} else if strings.Contains(err.Error(), constants.ServiceNotFoundError) {
			errorMsg := models.ErrorMessage{Key: constants.Services, ErrorMessage: constants.ServiceNotFoundError}
			errorResponse := commonModels.ErrorAPIResponse{
				Message: errorMsg,
				Error:   constants.VisitRecordCreationFailedError,
			}
			ctx.IndentedJSON(http.StatusNotFound, errorResponse)
			return
		} else if strings.Contains(err.Error(), constants.SomeServicesNotAvailableError) {
			errorMsg := models.ErrorMessage{Key: constants.Services, ErrorMessage: err.Error()}
			errorResponse := commonModels.ErrorAPIResponse{
				Message: errorMsg,
				Error:   constants.VisitRecordCreationFailedError,
			}
			ctx.IndentedJSON(http.StatusBadRequest, errorResponse)
			return
		}

		errorResponse := commonModels.ErrorAPIResponse{
			Error: constants.VisitRecordCreationFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, visitModels.BFFAddVisitRecordsResponse{
		Status: constants.VisitRecordCreationSuccess,
	})

}
