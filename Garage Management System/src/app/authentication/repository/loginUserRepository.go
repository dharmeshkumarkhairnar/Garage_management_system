package repository

import (
	"context"
	"errors"
	"fmt"
	"garage_management_system/src/app/authentication/commons/constants"
	genericModels "garage_management_system/src/models"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type loginUserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewLoginUserRepository(db *gorm.DB, logger *logrus.Logger) *loginUserRepository {
	return &loginUserRepository{
		db:     db,
		logger: logger,
	}
}

type LoginUserRepository interface {
	LoginUser(ctx context.Context, userEmail string) (genericModels.Users, error)
}

func (repo *loginUserRepository) LoginUser(ctx context.Context, userEmail string) (genericModels.Users, error) {
	var user genericModels.Users

	err := repo.db.WithContext(ctx).Table(constants.UsersTableName).Where(constants.EmailPlaceholder, userEmail).First(&user).Error
	if err != nil {
		fmt.Println("ERROR: ", err)
		if strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error()) {
			return user, errors.New(constants.UserNotFoundError)
		}

		return user, err
	}
	return user, nil
}
