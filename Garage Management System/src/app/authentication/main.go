package main

import (
	"garage_management_system/src/app/authentication/router"
	"garage_management_system/src/utils/database"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	db := database.GetDB()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	
	startRouter(db.DB, logger)
}

func startRouter(db *gorm.DB, logger *logrus.Logger) {
	router := router.GetRouter(db, logger)
	router.Run(":8080")
}
