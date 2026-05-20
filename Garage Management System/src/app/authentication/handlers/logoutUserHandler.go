package handlers

import (
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogoutUserHandler struct {
	service *business.LogoutUserService
}

func NewLogoutUserHandler(service *business.LogoutUserService) *LogoutUserHandler {
	return &LogoutUserHandler{
		service: service,
	}
}

// HandleLoginUSer handles the user/admin logout request.
// @Summary Logout registered user or admin
// @Description Logouts user by deleting token from redis
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.BFFLogoutUserResponse "Logout successful"
// @Failure 500 {object} models.BFFLogoutUserResponse "Internal Server Error"
// @Router /api/auth/logout [post]
func (controller *LogoutUserHandler) HandleLogoutUSer(ctx *gin.Context) {

	logger := logrus.New()
	tokenString := ctx.GetString(constants.Token)

	err := controller.service.LogoutUser(ctx, logger, tokenString)
	if err != nil {

		ctx.IndentedJSON(http.StatusInternalServerError, models.ErrorAPIResponse{
			Message: models.ErrorMessage{Key: constants.Redis, ErrorMessage: constants.RedisOperationError},
			Error:   constants.LogoutFailedError,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, models.BFFLogoutUserResponse{
		Message: " logged out successfully",
	})
}
