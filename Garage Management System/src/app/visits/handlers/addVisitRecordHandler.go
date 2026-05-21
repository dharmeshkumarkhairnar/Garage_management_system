package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/commons/constants"
	visitModels "garage_management_system/src/app/visits/models"
	commonModels "garage_management_system/src/models"
	"garage_management_system/src/utils/validations"
	"net/http"

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

// HandlerCreaterVehicle handles the Vehicle creation request.
// @Summary Create a new Vehicle
// @Description Handles Vehicle registration by validating input and storing Vehicle details
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param request body models.BFFAddVisitRecordsRequest true "Vehicle Registration Request"
// @Success 201 {object} models.BFFAddVisitRecordsResponse "User created successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 404 {object} models.ErrorAPIResponse "User not found"
// @Failure 409 {object} models.ErrorAPIResponse "Duplicate value in request"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/vehicles/create-vehicle [post]
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
		// if strings.Contains(err.Error(), constants.VehicleNumberPlateAlreadyExistsError) {
		// 	errorMsg := models.ErrorMessage{Key: "number plate", ErrorMessage: constants.VehicleNumberPlateAlreadyExistsError}
		// 	errorResponse := commonModels.ErrorAPIResponse{
		// 		Message: errorMsg,
		// 		Error:   constants.VehicleCreationFailedError,
		// 	}
		// 	ctx.IndentedJSON(http.StatusConflict, errorResponse)
		// 	return
		// } else if strings.Contains(err.Error(), constants.UserNotFoundError) {
		// 	errorMsg := models.ErrorMessage{Key: "user", ErrorMessage: constants.UserNotFoundError}
		// 	errorResponse := commonModels.ErrorAPIResponse{
		// 		Message: errorMsg,
		// 		Error:   constants.VehicleCreationFailedError,
		// 	}
		// 	ctx.IndentedJSON(http.StatusNotFound, errorResponse)
		// 	return
		// }

		errorResponse := commonModels.ErrorAPIResponse{
			Error: constants.VehicleCreationFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, visitModels.BFFAddVisitRecordsResponse{
		Status: constants.VisitRecordCreationSuccess,
	})

}
