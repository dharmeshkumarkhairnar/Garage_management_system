package handlers

import (
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

func (controller DeleteServiceHandler) HandleDeleteService(ctx *gin.Context) {
	serviceName := strings.ToLower(ctx.Param("serviceName"))

	err := controller.service.DeleteService(ctx.Request.Context(), serviceName)
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

	ctx.IndentedJSON(http.StatusOK, models.BFFDeleteerviceResponse{
		Message: "service deleted successfully",
	})
}
