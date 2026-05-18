package handlers

import (
	"encoding/json"
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	genericModels "garage_management_system/src/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateCustomer struct {
	service *business.CreateCustomer
}

func NewCreateCustomer(service *business.CreateCustomer) *CreateCustomer {

	return &CreateCustomer{
		service: service,
	}
}

func (controller *CreateCustomer) HandleCreateCustomer(ctx *gin.Context) {
	var bffCreateCustomerRequest models.BFFCreateCustomerRequest

	if err := ctx.ShouldBind(&bffCreateCustomerRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: constants.ErrUnexpectedValue,
		}

		ctx.IndentedJSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
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

	err := controller.service.CreateNewCustomer(ctx.Request.Context(), bffCreateCustomerRequest)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrDuplicateEntry) {	
			errorResponse := genericModels.ErrorAPIResponse{
				Message: genericModels.ErrorMessage{
					Key:          strings.Split(err.Error(), constants.ErrDuplicateEntry)[0],
					ErrorMessage: constants.ErrUserAlreadyExists,
				},
				Error: constants.ErrConflict,
			}
			ctx.IndentedJSON(http.StatusConflict, errorResponse)
			return
		}

		errorResponse := genericModels.ErrorAPIResponse{
			Error: constants.ErrUserCreationFailed,
		}

		ctx.IndentedJSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, constants.UserCreationSuccessMsg)
}
