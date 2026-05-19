package router

import (
	"garage_management_system/src/app/visits/middleware"
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/constants"
	"garage_management_system/src/app/visits/handlers"
	"garage_management_system/src/app/visits/repository"

	"garage_management_system/src/utils/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"

	"garage_management_system/src/app/visits/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	gdb := database.GetDB().DB

	docs.SwaggerInfo.Title = "Garage management system"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:    []string{"Authorization", "Content-type", "Origin"},
	}))

	createVehicleRepository := repository.NewCreateVehicleRepository(gdb)
	createVehicleService := business.NewCreateVehicleService(createVehicleRepository)
	createVehicleHandler := handlers.NewCreateVehicleHandler(createVehicleService)

	Group := router.Group(constants.RoutePrefix)
	{
		Group.POST(constants.CreateVehicle,middleware.VisitMiddleware(), createVehicleHandler.CreaterVehicle)
	}

	return router
}
