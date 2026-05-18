package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/constants"
	visitModels "garage_management_system/src/app/visits/models"
	commonModels "garage_management_system/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreaterVehicleHandler struct {
	service *business.CreateVehicleService
}

func NewCreateVehicleHandler(service *business.CreateVehicleService) *CreaterVehicleHandler {
	return &CreaterVehicleHandler{
		service: service,
	}
}

// HandlerCreaterVehicle handles the Vehicle creation request.
// @Summary Create a new Vehicle
// @Description Handles Vehicle registration by validating input and storing Vehicle details
// @Tags Vehicles
// @Accept json
// @Produce json
// @Param request body visitModels.BFFCreateVehicleRequest true "Vehicle Registration Request"
// @Success 201 {object} visitModels.BFFCreateVehicleResponse "User created successfully"
// @Failure 400 {object} commonModels.ErrorAPIResponse "Invalid input payload"
// @Failure 500 {object} commonModels.ErrorAPIResponse "Internal Server Error"
// @Router /api/vehicles/create-vehicle [post]
func (controller *CreaterVehicleHandler) CreaterVehicle(ctx *gin.Context) {

	var bffCreateVehicleRequest visitModels.BFFCreateVehicleRequest

	if err := ctx.ShouldBind(&bffCreateVehicleRequest); err != nil {
		errorMsgs := commonModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, commonModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	// if err := validations.GetBFFValidator().Struct(&bffCreateUserRequest); err != nil {
	// 	validationErros, _ := validations.FormatValidationErrors(err)
	// 	ctx.IndentedJSON(http.StatusBadRequest, validationErros)
	// 	return
	// }

	err := controller.service.CreateVehicle(ctx, ctx.Request.Context(), bffCreateVehicleRequest)
	if err != nil {
		errorResponse := commonModels.ErrorAPIResponse{
			Error: constants.VehicleCreationFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, constants.VehicleCreationSuccess)

}
