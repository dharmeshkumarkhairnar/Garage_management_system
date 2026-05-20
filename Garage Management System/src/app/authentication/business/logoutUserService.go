package business

import (
	"context"
	"errors"
	"fmt"
	"garage_management_system/src/app/authentication/commons/constants"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type LogoutUserService struct {
	redisClient *redis.Client
}

func NewLogoutUserService(redisClient *redis.Client) *LogoutUserService {
	return &LogoutUserService{
		redisClient: redisClient,
	}
}

func (service *LogoutUserService) LogoutUser(ctx context.Context, logger *logrus.Logger, tokenString string) error {

	cacheKey := fmt.Sprintf(constants.RedisTokenCacheKey, tokenString)

	err := service.redisClient.Expire(ctx, cacheKey, 0).Err()

	if err != nil {
		return errors.New(constants.RedisOperationError)
	}

	logger.Info(constants.RedisTokenRemoved)

	return nil
}
