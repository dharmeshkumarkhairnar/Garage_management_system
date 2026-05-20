package business

import (
	"context"
	"fmt"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/models"
	"garage_management_system/src/app/authentication/repository"
	"garage_management_system/src/utils"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LoginUserService struct {
	db          *gorm.DB
	redisClient *redis.Client
	repo        repository.LoginUserRepository
}

func NewLoginUserService(db *gorm.DB, redisClient *redis.Client, repo repository.LoginUserRepository) *LoginUserService {
	return &LoginUserService{
		db:          db,
		redisClient: redisClient,
		repo:        repo,
	}
}

func (service *LoginUserService) LoginUser(ctx context.Context, bffLoginUserRequest models.BFFLoginUserRequest) (string, error) {

	user, err := service.repo.LoginUser(ctx, bffLoginUserRequest.Email)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	if strings.ToLower(bffLoginUserRequest.Role) != user.Role {
		return "", fmt.Errorf(constants.RoleMismatchError)
	}

	if !utils.CompareHashPassword(user.Password, bffLoginUserRequest.Password) {
		return "", fmt.Errorf(constants.IncorrectPasswordError)
	}

	if service.redisClient == nil {
		return "", fmt.Errorf("redis client not initialized")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", fmt.Errorf("")
	}

	TokenCacheKey := fmt.Sprintf(constants.RedisTokenCacheKey, token)
	if err := service.redisClient.Set(ctx, TokenCacheKey, user.ID, 24*time.Hour).Err(); err != nil {
		return "", err
	}
	fmt.Println("CACHE KEY STORED:", TokenCacheKey)

	fmt.Println("TOKEN:", token)
	return token, nil
}
