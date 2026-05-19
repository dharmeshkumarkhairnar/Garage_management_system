package router

import (
	"garage_management_system/src/utils/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"

	"garage_management_system/src/app/assets/business"
	"garage_management_system/src/app/assets/constants"
	"garage_management_system/src/app/assets/docs"
	"garage_management_system/src/app/assets/handlers"
	"garage_management_system/src/app/assets/middleware"
	"garage_management_system/src/app/assets/repository"

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

	addMechanicRepository := repository.NewAddMechanicRepository(gdb)
	addMechanicService := business.NewAddMechanicService(addMechanicRepository)
	addMechanicHandler := handlers.NewAddMechanicHandler(addMechanicService)

	deleteMechanicRepository := repository.NewDeleteMechanicRepository(gdb)
	deleteMechanicService := business.NewDeleteMechanicService(deleteMechanicRepository)
	deleteMechanicHandler := handlers.NewDeleteMechanicHandler(deleteMechanicService)

	Group := router.Group(constants.RoutePrefix)
	{
		Group.POST(constants.AddMechanic, middleware.AssetMiddleware(), addMechanicHandler.AddMechanic)
		Group.POST(constants.DeleteMechanic, middleware.AssetMiddleware(), deleteMechanicHandler.DeleteMechanic)
	}

	return router
}
