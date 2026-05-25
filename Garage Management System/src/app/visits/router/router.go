package router

import (
	"garage_management_system/src/app/visits/business"
	"garage_management_system/src/app/visits/commons/constants"
	"garage_management_system/src/app/visits/handlers"
	"garage_management_system/src/app/visits/middleware"
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

	addVisitRecordsRepository := repository.NewAddVisitRecordsRepository(gdb)
	addVisitRecordsService := business.NewAddVisitRecordService(addVisitRecordsRepository)
	addVisitRecordsHandler := handlers.NewAddVisitRecordHandler(addVisitRecordsService)

	getBillRepository := repository.NewGenerateBillRepository(gdb)
	getBillService := business.NewGenerateBillervice(getBillRepository)
	getBillHandler := handlers.NewGenerateBillHandler(getBillService)

	VGroup := router.Group(constants.VehiclesRoutePrefix)
	{
		VGroup.POST(constants.CreateVehicle, middleware.VisitMiddleware(), createVehicleHandler.CreateVehicle)
	}

	VisitsGroup := router.Group(constants.VisitsRoutePrefix)
	{
		VisitsGroup.POST(constants.AddVisitRecord, middleware.VisitMiddleware(), addVisitRecordsHandler.AddVisitRecord)
		VisitsGroup.POST(constants.GenerateBill, getBillHandler.GenerateBill)
	}

	return router
}
