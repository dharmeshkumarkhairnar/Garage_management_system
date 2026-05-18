package router

import (
	"garage_management_system/src/app/authentication/business"
	"garage_management_system/src/app/authentication/handlers"
	"garage_management_system/src/app/authentication/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetRouter(db *gorm.DB, logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// docs.SwaggerInfo.Title = "garage management system"
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:    []string{"Authorization", "Content-type", "Origin"},
	}))

	createCustomerRepo := repository.NewCreateCustomer(db, logger)
	createCustomerSvc := business.NewCreateUserService(createCustomerRepo, db, logger)
	createCustomerHandler := handlers.NewCreateCustomer(createCustomerSvc)

	router.POST("/api/auth/register/customer", createCustomerHandler.HandleCreateCustomer)
	return router
}
