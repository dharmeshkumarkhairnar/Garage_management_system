package router

import (
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/commons/constants"
	"garage_management_system/src/app/authentication/handlers"
	"garage_management_system/src/app/authentication/middleware"
	"garage_management_system/src/app/authentication/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// "github.com/swaggo/swag/example/basic/docs"
	"garage_management_system/src/app/authentication/docs"

	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB, logger *logrus.Logger, redisClient *redis.Client) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	docs.SwaggerInfo.Title = "garage management system"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:    []string{"Authorization", "Content-type", "Origin"},
	}))

	createCustomerRepo := repository.NewCreateCustomerRepository(db, logger)
	createCustomerSvc := business.NewCreateUserService(createCustomerRepo, db, logger)
	createCustomerHandler := handlers.NewCreateCustomerHandler(createCustomerSvc)

	loginUserRepository := repository.NewLoginUserRepository(db, logger)
	loginUserService := business.NewLoginUserService(db, redisClient, loginUserRepository)
	loginUserHandler := handlers.NewLoginUserHandler(loginUserService)

	logoutUserService := business.NewLogoutUserService(redisClient)
	logoutUserHandler := handlers.NewLogoutUserHandler(logoutUserService)
	AuthGroup := router.Group(constants.AuthRoutePrefix)
	{
		AuthGroup.POST(constants.RegisterRoute, createCustomerHandler.HandleCreateCustomer)
		AuthGroup.POST(constants.LoginUserRoute, loginUserHandler.HandleLoginUSer)
		AuthGroup.POST(constants.LogoutUserRoute,middleware.AuthMiddleware(), logoutUserHandler.HandleLogoutUSer)
	}
	return router
}
