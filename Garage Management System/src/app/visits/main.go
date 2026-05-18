package main

import (
	"garage_management_system/src/app/visits/router"
	"garage_management_system/src/utils/database"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// err = redis.InitRedis()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize redis: %v", err)
	// }

	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info("")
	router.Run(":8080")
}
