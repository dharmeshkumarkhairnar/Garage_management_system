package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/assets/business"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genModels "garage_management_system/src/models"
	"garage_management_system/src/utils/validations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddMechanicHandler struct {
	service *business.AddMechanicService
}

func NewAddMechanicHandler(service *business.AddMechanicService) *AddMechanicHandler {
	return &AddMechanicHandler{
		service: service,
	}
}

// HandlerAddMechanic handles the mechanic add request.
// @Summary Adds new mechanic
// @Description Handles mechanic add request by validating input and storing mechanics details
// @Tags Mechanics
// @Accept json
// @Produce json
// @Param request body models.BFFAddMechanicRequest true "mechanic add Request"
// @Success 201 {object} models.BFFAddMechanicResponse "mechanic added successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 409 {object} models.ErrorAPIResponse "Duplicate value in request"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server error"
// @Router /api/mechanics/add [post]
func (controller *AddMechanicHandler) AddMechanic(ctx *gin.Context) {

	var bffAddMechanicRequest models.BFFAddMechanicRequest

	if err := ctx.ShouldBind(&bffAddMechanicRequest); err != nil {
		errorMsgs := genModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, genModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	if err := validations.GetBFFValidator().Struct(&bffAddMechanicRequest); err != nil {
		validationErros, _ := validations.FormatValidationErrors(err)
		ctx.IndentedJSON(http.StatusBadRequest, validationErros)
		return
	}

	err := controller.service.AddMechanic(ctx, ctx.Request.Context(), bffAddMechanicRequest)
	if err != nil {
		if strings.Contains(err.Error(), constants.DuplicateAadharNumberError) {
			errorMsg := genModels.ErrorMessage{Key: constants.FieldMechanicAadharNumber, ErrorMessage: constants.DuplicateAadharNumberError}
			errorResponse := genModels.ErrorAPIResponse{
				Message: errorMsg,
				Error:   constants.MechanicAdditionFailedError,
			}
			ctx.IndentedJSON(http.StatusConflict, errorResponse)
			return
		}

		errorResponse := genModels.ErrorAPIResponse{
			Error: constants.MechanicAdditionFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, models.BFFAddMechanicResponse{
		Status: constants.MechanicAddedSuccessfully,
	})

}
