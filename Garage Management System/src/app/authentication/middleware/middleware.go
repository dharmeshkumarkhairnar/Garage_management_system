package middleware

import (
	"fmt"
	"garage_management_system/src/app/assets/commons/constants"
	"garage_management_system/src/app/assets/models"
	genConst "garage_management_system/src/constants"
	genModels "garage_management_system/src/models"
	"garage_management_system/src/utils"
	"garage_management_system/src/utils/redis"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		defer func() {
			log.Printf("completed request in %d ms", time.Since(start).Milliseconds())
		}()

		authHeader := ctx.GetHeader(constants.Authorization)

		if len(authHeader) == 0 {
			errorMsg := genModels.ErrorMessage{Key: constants.Header, ErrorMessage: constants.HeaderIsMissingError}
			response := genModels.ErrorAPIResponse{Message: errorMsg, Error: constants.OperationFailed}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, constants.Bearer)

		redisClient := redis.GetRedisClient()

		cacheKey := fmt.Sprintf(constants.ActiveToken, tokenString)

		count, err := redisClient.Exists(ctx, cacheKey).Result()

		if err != nil {
			errorMsg := genModels.ErrorMessage{Key: constants.Redis, ErrorMessage: constants.RedisOperationError}
			response := genModels.ErrorAPIResponse{Message: errorMsg, Error: constants.OperationFailed}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		if count <= 0 {
			errorMsg := genModels.ErrorMessage{Key: constants.User, ErrorMessage: constants.UserLoggedOutError}
			response := genModels.ErrorAPIResponse{Message: errorMsg, Error: constants.OperationFailed}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := utils.ParseToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{
					Key:          constants.Token,
					ErrorMessage: err.Error(),
				},
				Error: constants.OperationFailed,
			})
			return
		}

		claims, err := utils.VerifyToken(token)

		if err != nil {
			if strings.Contains(err.Error(), genConst.ClaimMappingFailedError) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
					Message: models.ErrorMessage{
						Key:          constants.Token,
						ErrorMessage: genConst.ClaimMappingFailedError,
					},
					Error: constants.OperationFailed,
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{
					Key:          constants.Token,
					ErrorMessage: constants.InvalidTokenError,
				},
				Error: constants.OperationFailed,
			})
			return
		}

		userId, ok := claims[constants.Subject].(float64)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{
					Key:          constants.Token,
					ErrorMessage: constants.MissingUserIDError,
				},
				Error: constants.OperationFailed,
			})
			return
		}

		userRole, ok := claims[constants.UserRole].(string)

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorAPIResponse{
				Message: models.ErrorMessage{
					Key:          constants.Token,
					ErrorMessage: constants.MissingUserRoleError,
				},
				Error: constants.OperationFailed,
			})
			return
		}


		ctx.Set(constants.UserId, uint64(userId))
		ctx.Set(constants.Token, tokenString)
		ctx.Set(constants.UserRole, userRole)

		ctx.Next()
	}
}
