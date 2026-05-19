package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/assets/business"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genericModels "garage_management_system/src/models"
	"garage_management_system/src/utils/validations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AddNewServiceHandler struct {
	service *business.AddNewServiceService
}

func NewAddNewServiceHandler(service *business.AddNewServiceService) *AddNewServiceHandler {
	return &AddNewServiceHandler{
		service: service,
	}
}

func (controller AddNewServiceHandler) HandleAddNewService(ctx *gin.Context) {
	var bffAddNewServiceRequest models.BFFAddNewServiceRequest

	if err := ctx.ShouldBind(&bffAddNewServiceRequest); err != nil {
		errMsgs := genericModels.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: constants.ErrUnexpectedValue,
		}

		ctx.IndentedJSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
			Message: errMsgs,
			Error:   constants.ErrInvalidPayload,
		})
		return
	}

	if err := validations.GetBFFValidator().Struct(&bffAddNewServiceRequest); err != nil {
		validationErros, _ := validations.FormatValidationErrors(err)
		ctx.IndentedJSON(http.StatusBadRequest, validationErros)
		return
	}

	err := controller.service.AddNewService(ctx, bffAddNewServiceRequest)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrServiceAlreadyExists) {
			errorResponse := models.ErrorAPIResponse{
				Message: models.ErrorMessage{
					Key: "service", ErrorMessage: "service already exists",
				},
				Error: constants.ErrConflict,
			}

			ctx.IndentedJSON(http.StatusConflict, errorResponse)
			return
		}

		errorResponse := genericModels.ErrorMessage{
			Key: "service", ErrorMessage: "failed to add new service",
		}
		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, models.BFFAddNewServiceResponse{Message: "added new service successfully"})
}
