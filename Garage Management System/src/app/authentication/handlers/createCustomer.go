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

// HandlerCreaterCustomer handles the customer creation request.
// @Summary Create a new customer
// @Description Handles customer registration by validating input and storing user details
// @Tags Customer
// @Accept json
// @Produce json
// @Param request body models.BFFCreateCustomerRequest true "Customer Registration Request"
// @Success 201 {string} string "Customer created successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 409 {object} models.ErrorAPIResponse "User already exists"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/auth/register/customer [post]
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
