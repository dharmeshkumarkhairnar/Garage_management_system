package handlers

import (
	"encoding/json"
	"fmt"
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	genericModels "garage_management_system/src/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateCustomerHandler struct {
	service *business.CreateCustomerService
}

func NewCreateCustomerHandler(service *business.CreateCustomerService) *CreateCustomerHandler {

	return &CreateCustomerHandler{
		service: service,
	}
}

// HandlerCreaterCustomer handles the customer creation request.
// @Summary Create a new customer
// @Description Handles customer registration by validating input and storing user details
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.BFFCreateCustomerRequest true "Customer Registration Request"
// @Success 201 {string} string "Customer created successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input payload"
// @Failure 409 {object} models.ErrorAPIResponse "User already exists"
// @Failure 500 {object} models.ErrorAPIResponse "Internal Server Error"
// @Router /api/auth/register [post]
func (controller *CreateCustomerHandler) HandleCreateCustomer(ctx *gin.Context) {
	var bffCreateCustomerRequest models.BFFCreateCustomerRequest

	if err := ctx.ShouldBind(&bffCreateCustomerRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: constants.UnexpectedValueError,
		}

		ctx.IndentedJSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
			Message: errorMsgs,
			Error:   constants.InvalidPayloadError,
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
		fmt.Println("ERROR:", err)
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

	ctx.IndentedJSON(http.StatusCreated, models.BFFCreateCustomerResponse{
		Message: "customer created successfully",
	})
}
