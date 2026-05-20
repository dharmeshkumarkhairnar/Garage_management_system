package utils

import (
	"errors"
	constant "garage_management_system/src/constants"
	"garage_management_system/src/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey *models.JWT

func GenerateToken(userID uint64, role string) (string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(secretKey.AccessSecretKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func ParseToken(tokenstring string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constant.WrongSigningAlgorithmError)
		}

		return []byte(secretKey.AccessSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func VerifyToken(token *jwt.Token) (jwt.MapClaims, error) {

	if !token.Valid {
		return nil, errors.New(constant.TokenExpiredError)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New(constant.ClaimMappingFailedError)
	}

	return claims, nil
}

func InitJWTConfig() {
	secretKey = &models.JWT{AccessSecretKey: "Apr/meTe4sxpBwxb36ISTRNnHc4y+Y34KjQ/ntwB1Kw=", RefreshSecretKey: "uFEnDER3W78vCqGX86jsbfDzrGsutR8m06+d1ya3JK0="}
}
