package handlers

// import (
// 	"encoding/json"
// 	"garage_management_system/src/app/assets/business"
// 	"garage_management_system/src/app/assets/models"
// 	"garage_management_system/src/app/assets/visits/constants"
// 	genModels "garage_management_system/src/models"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// type AddMechanicHandler struct {
// 	service *business.AddMechanicService
// }

// func NewAddMechanicHandler(service *business.AddMechanicService) *AddMechanicHandler {
// 	return &AddMechanicHandler{
// 		service: service,
// 	}
// }

// // HandlerAddMechanic handles the mechanic add request.
// // @Summary Create a new Vehicle
// // @Description Handles Vehicle registration by validating input and storing Vehicle details
// // @Tags Vehicles
// // @Accept json
// // @Produce json
// // @Param request body models.BFFCreateVehicleRequest true "Vehicle Registration Request"
// // @Success 201 {object} models.BFFCreateVehicleResponse "User created successfully"
// // @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// // @Failure 404 {object} models.ErrorAPIResponse "User not found"
// // @Failure 409 {object} models.ErrorAPIResponse "Duplicate value in request"
// // @Failure 500 {object} models.ErrorAPIResponse "Internal Servebusiness
// // @Router /api/vehicles/create-vehicle [post]
// func (controller *AddMechanicHandler) AddMechanic(ctx *gin.Context) {

// 	var bffAddMechanicRequest models.BFFAddMechanicRequest

// 	if err := ctx.ShouldBind(&bffAddMechanicRequest); err != nil {
// 		errorMsgs := genModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
// 		ctx.IndentedJSON(http.StatusBadRequest, genModels.ErrorAPIResponse{
// 			Message: errorMsgs,
// 			Error:   constants.ErrInvalidPayload,
// 		})
// 		return
// 	}

// 	// if err := validations.GetBFFValidator().Struct(&bffCreateUserRequest); err != nil {
// 	// 	validationErros, _ := validations.FormatValidationErrors(err)
// 	// 	ctx.IndentedJSON(http.StatusBadRequest, validationErros)
// 	// 	return
// 	// }

// 	err := controller.service.AddMechanic(ctx, ctx.Request.Context(), bffAddMechanicRequest)
// 	if err != nil {
// 		// if strings.Contains(err.Error(), constants.VehicleNumberPlateAlreadyExistsError) {
// 		// 	errorMsg := models.ErrorMessage{Key: "number plate", ErrorMessage: constants.VehicleNumberPlateAlreadyExistsError}
// 		// 	errorResponse := genModels.ErrorAPIResponse{
// 		// 		Message: errorMsg,
// 		// 		Error:   constants.VehicleCreationFailedError,
// 		// 	}
// 		// 	ctx.IndentedJSON(http.StatusConflict, errorResponse)
// 		// 	return
// 		// } else if strings.Contains(err.Error(), constants.UserNotFoundError) {
// 		// 	errorMsg := models.ErrorMessage{Key: "user", ErrorMessage: constants.UserNotFoundError}
// 		// 	errorResponse := commonModels.ErrorAPIResponse{
// 		// 		Message: errorMsg,
// 		// 		Error:   constants.VehicleCreationFailedError,
// 		// 	}
// 		// 	ctx.IndentedJSON(http.StatusNotFound, errorResponse)
// 		// 	return
// 		// }

// 		errorResponse := commonModels.ErrorAPIResponse{
// 			Error: constants.VehicleCreationFailedError,
// 		}
// 		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
// 		return
// 	}

// 	ctx.IndentedJSON(http.StatusCreated, constants.VehicleCreationSuccess)

// }
