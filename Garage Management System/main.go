package main

import (
	model "garage_management_system/src/models"
	"garage_management_system/src/utils/database"
	"garage_management_system/src/utils/redis"
	"log"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	db := database.GetDB().DB

	if err := db.AutoMigrate(&model.Customers{}, &model.Vehicles{}, &model.Mechanics{}, &model.ServiceMaster{}, &model.VisitRecords{}, &model.VisitServices{}); err != nil {
		log.Fatalf("error in automigrating schema: %s", err.Error())
	}

	log.Print("Database migrated successfully")

	err = redis.InitRedis()
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}

}
