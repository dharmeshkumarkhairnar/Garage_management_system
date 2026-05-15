package database

import (
	"errors"
	"fmt"
	constant "garage_management_system/src/constants"
	"garage_management_system/src/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *model.Database
var once sync.Once

// InitDB opens a PostgreSQL connection via GORM using DATABASE_URL
// InitDB initializes the database only once
func InitDB() error {

	var initErr error

	once.Do(func() {
		dsn := fmt.Sprintf(constant.DSNString, "localhost", "5432", "garage_management_system", "postgres", "mysql@1715", "Asia/Kolkata")
		if dsn == "" {
			initErr = errors.New("DATABASE_URL is not set")
			return
		}

		gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			initErr = fmt.Errorf("open database: %w", err)
			return
		}

		db = &model.Database{
			DB: gdb,
		}
	})

	return initErr
}

// GetDB returns the shared model.Database. Call InitDB first.
func GetDB() *model.Database {
	return db
}
