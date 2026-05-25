package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/models"
	commonModels "garage_management_system/src/models"
	"garage_management_system/src/utils/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerateBillHandler struct {
	service *business.GenerateBillService
}

func NewGenerateBillHandler(service *business.GenerateBillService) *GenerateBillHandler {
	return &GenerateBillHandler{
		service: service,
	}
}

// HandlerGenerateBill handles the visitors bill request.
// @Summary Generate Bill for the Visit
// @Description Handles generation of bill for the respective visit
// @Tags Visits
// @Accept json
// @Produce json
// @Param request body models.BFFGenerateBillRequest true "Bill Generation Request"
// @Success 201 {object} models.BFFGenerateBillResponse "Bill Generated Successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 404 {object} models.ErrorAPIResponse "User not found"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/visits/get-bill [post]
func (controller *GenerateBillHandler) GenerateBill(ctx *gin.Context) {

	var bffGenerateBillRequest models.BFFGenerateBillRequest

	if err := ctx.ShouldBind(&bffGenerateBillRequest); err != nil {
		errorMsgs := commonModels.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, commonModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	if err := validations.GetBFFValidator().Struct(&bffGenerateBillRequest); err != nil {
		validationErros, _ := validations.FormatValidationErrors(err)
		ctx.IndentedJSON(http.StatusBadRequest, validationErros)
		return
	}

	billAmount, err := controller.service.GenerateBill(ctx, ctx.Request.Context(), bffGenerateBillRequest)
	if err != nil {
		errorResponse := commonModels.ErrorAPIResponse{
			Error: constants.BillGenerationFailedError,
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, models.BFFGenerateBillResponse{
		BillAmount: billAmount,
		Status:     constants.BillGenerationSuccess,
	})

}
