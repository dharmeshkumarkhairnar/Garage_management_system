package main

import (
	model "garage-system/models"
	"garage-system/router"
	"garage-system/utils/database"
	"garage-system/utils/redis"
	"log"

	"github.com/sirupsen/logrus"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB().DB

	if err := db.AutoMigrate(&model.Customers{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	if err := db.AutoMigrate(&model.Vehicles{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	if err := db.AutoMigrate(&model.Mechanics{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	if err := db.AutoMigrate(&model.ServiceRecords{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	log.Print("Database migrated successfully")

	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

	startRouter()
}

func startRouter() {
	logger := logrus.New()
	router := router.GetRouter()
	logger.Info("")
	router.Run(":8080")
}
