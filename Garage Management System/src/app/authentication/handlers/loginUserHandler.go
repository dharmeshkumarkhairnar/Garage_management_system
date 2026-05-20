package handlers

import (
	"encoding/json"
	"fmt"
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type LoginUserHandler struct {
	service *business.LoginUserService
}

func NewLoginUserHandler(service *business.LoginUserService) *LoginUserHandler {
	return &LoginUserHandler{
		service: service,
	}
}

// HandleLoginUSer handles the user/admin login request.
// @Summary Login registered user or admin
// @Description Authenticates user and returns JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param request body models.BFFLoginUserRequest true "User Sign In Request"
// @Success 200 {object} models.BFFLoginUserResponse "Login successful"
// @Failure 400 {object} models.BFFLoginUserResponse "Invalid input payload"
// @Failure 401 {object} models.BFFLoginUserResponse "Invalid credentials"
// @Failure 404 {object} models.BFFLoginUserResponse "User does not exist"
// @Failure 500 {object} models.BFFLoginUserResponse "Internal Server Error"
// @Router /api/auth/login [post]
func (controller *LoginUserHandler) HandleLoginUSer(ctx *gin.Context) {
	var req models.BFFLoginUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		errMsgs := models.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: constants.UnexpectedValueError}

		ctx.IndentedJSON(http.StatusBadRequest, models.ErrorAPIResponse{
			Message: errMsgs,
			Error:   constants.LoginReuestFailedError,
		})
		return
	}
	token, err := controller.service.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		fmt.Println("Handler ERROR: ", err)
		if strings.Contains(err.Error(), constants.UserNotFoundError) {
			ctx.IndentedJSON(http.StatusNotFound, models.ErrorAPIResponse{
				Message: models.ErrorMessage{Key: "user", ErrorMessage: constants.UnauthorizedRequestError},
				Error:   constants.LoginReuestFailedError,
			})
			return
		}

		if strings.Contains(err.Error(), constants.IncorrectPasswordError) {
			ctx.IndentedJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{Key: "request", ErrorMessage: constants.IncorrectPasswordError},
				Error:   constants.UnauthorizedRequestError,
			})
			return
		}

		if strings.Contains(err.Error(), constants.RoleMismatchError) {
			ctx.IndentedJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{Key: "request", ErrorMessage: constants.UnauthorizedRequestError},
				Error:   constants.LoginReuestFailedError,
			})
			return
		}

		ctx.IndentedJSON(http.StatusInternalServerError, models.ErrorAPIResponse{
			Message: models.ErrorMessage{Key: "server", ErrorMessage: constants.InternalServerError},
			Error:   constants.LoginReuestFailedError,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, models.BFFLoginUserResponse{
		Token:   token,
		Message: fmt.Sprintf("%s loggedin successfully", strings.ToLower(req.Role)),
	})
}
