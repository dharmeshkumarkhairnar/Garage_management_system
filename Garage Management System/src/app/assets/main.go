package main

import (
	"garage_management_system/src/app/assets/router"
	"garage_management_system/src/utils"
	"garage_management_system/src/utils/database"
	"garage_management_system/src/utils/redis"
	"log"

	"github.com/sirupsen/logrus"
)

// @title Service Operations API
// @version 1.0
// @description Services API for Garage Management System
// @query.collection.format multi
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @x-extension-openapi {"example": "value on a json format"}
func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	//Initialised for JWT token
	utils.InitJWTConfig()

	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info("")
	router.Run(":8080")
}
