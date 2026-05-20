package handlers

import (
	"encoding/json"
	"fmt"
	"garage_management_system/src/app/assets/business"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genericModels "garage_management_system/src/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type DeleteServiceHandler struct {
	service *business.DeleteServiceService
}

func NewDeleteServiceHandler(service *business.DeleteServiceService) *DeleteServiceHandler {
	return &DeleteServiceHandler{
		service: service,
	}
}

// HandleDeleteService handles the requests for deleting existing services
// @Summary Delete an existing service
// @Description Deletes an existing service from the DB
// @Tags Services
// @Accept json
// @Produce json
// @Securirt BearerAuth
// @Param request body models.BFFDeleteServiceRequest true "Delete Service Request"
// @Success 200 {object} models.BFFDeleteServiceResponse "Deleted service successful"
// @Failure 400 {object} models.BFFDeleteServiceResponse "Invalid input payload"
// @Failure 401 {object} models.BFFDeleteServiceResponse "Unauthorized role"
// @Failure 500 {object} models.BFFDeleteServiceResponse "Internal Server Error"
// @Router /api/services/delete [delete]
func (controller DeleteServiceHandler) HandleDeleteService(ctx *gin.Context) {
	
	var bffDeleteServiceRequest models.BFFDeleteServiceRequest

	if err := ctx.ShouldBind(&bffDeleteServiceRequest); err != nil {
		errMsgs := models.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.ErrUnexpectedValue}
		ctx.IndentedJSON(http.StatusBadRequest, models.ErrorAPIResponse{
			Message: errMsgs, Error: constants.ErrInvalidPayload,
		})
		return
	}

	err := controller.service.DeleteService(ctx.Request.Context(), bffDeleteServiceRequest.Service)
	if err != nil {
		fmt.Println("ERROR: ", err)
		if strings.Contains(err.Error(), constants.ServiceDoesNotExist) {
			errMsgs := genericModels.ErrorMessage{
				Key:          constants.Service,
				ErrorMessage: constants.ServiceDoesNotExist,
			}

			ctx.IndentedJSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				Message: errMsgs,
				Error:   constants.FailedToDeleteService,
			})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
			Message: genericModels.ErrorMessage{
				Key: constants.Server, ErrorMessage: constants.ErrInternalServer},
			Error: constants.ErrInternalServer,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, models.BFFDeleteServiceResponse{
		Message: "service deleted successfully",
	})
}
