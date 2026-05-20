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

// HandleAddNewService handles the requests for adding new services
// @Summary Add a new service
// @Description Adds a new service to the DB
// @Tags Services
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.BFFAddNewServiceRequest true "Add Service Request"
// @Success 200 {object} models.BFFAddNewServiceResponse "Added service successful"
// @Failure 400 {object} models.BFFAddNewServiceResponse "Invalid input payload"
// @Failure 401 {object} models.BFFAddNewServiceResponse "Unauthorized role"
// @Failure 409 {object} models.BFFAddNewServiceResponse "Service already exists"
// @Failure 500 {object} models.BFFAddNewServiceResponse "Internal Server Error"
// @Router /api/services/add [post]
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
