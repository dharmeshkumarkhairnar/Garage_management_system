package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/assets/business"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genModels "garage_management_system/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteMechanicHandler struct {
	service *business.DeleteMechanicService
}

func NewDeleteMechanicHandler(service *business.DeleteMechanicService) *DeleteMechanicHandler {
	return &DeleteMechanicHandler{
		service: service,
	}
}

// HandlerDeleteMechanic handles the mechanic Delete request.
// @Summary Delete mechanic
// @Description Handles mechanic Delete request by validating input.
// @Tags Mechanics
// @Accept json
// @Produce json
// @Param request body models.BFFDeleteMechanicRequest true "mechanic Delete Request"
// @Success 200 {object} models.BFFDeleteMechanicResponse "mechanic Deleted successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server error"
// @Router /api/vehicles/delete-mechanic [post]
func (controller *DeleteMechanicHandler) DeleteMechanic(ctx *gin.Context) {

	var bffDeleteMechanicRequest models.BFFDeleteMechanicRequest

	if err := ctx.ShouldBind(&bffDeleteMechanicRequest); err != nil {
		errorMsgs := genModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, genModels.ErrorAPIResponse{
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

	err := controller.service.DeleteMechanic(ctx, ctx.Request.Context(), bffDeleteMechanicRequest)
	if err != nil {
		errorResponse := genModels.ErrorAPIResponse{
			Error: constants.MechanicDeletionFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, models.BFFAddMechanicResponse{
		Status: constants.MechanicDeletedSuccessfully,
	})

}
