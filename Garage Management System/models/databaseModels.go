package model

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	ID         uint64    `gorm:"column:id;primarykey" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Email      string    `gorm:"column:email;uniqueIndex" json:"email"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Created_at time.Time `gorm:"column:created_at" json:"created_at"`
}

type Vehicles struct {
	ID          uint64    `gorm:"column:id;primarykey" json:"id"`
	CustomerID  string    `gorm:"column:name" json:"name"`
	NumberPlate string    `gorm:"column:email;uniqueIndex" json:"email"`
	Model       string    `gorm:"column:phone" json:"phone"`
	Created_at  time.Time `gorm:"column:created_at" json:"created_at"`

	CustId Customers `gorm:"foreignkey:CustomerID;references:ID"`
}

type Mechanics struct {
	ID   uint64 `gorm:"column:id;primarykey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type ServiceRecords struct {
	ID         uint64    `gorm:"column:id;primarykey" json:"id"`
	VehicleID  uint64    `gorm:"column:vehicle_id" json:"vehicle_id"`
	MechanicID uint64    `gorm:"column:mechanic_id;uniqueIndex" json:"mechanic_id"`
	Service    string    `gorm:"column:service; not null" json:"service"`
	Price      float64 `gorm:"column:price" json:"price"`
	Date       time.Time `gorm:"column:created_at;type:date" json:"date"`

	Vehicle_ID  Vehicles  `gorm:"foreignkey:VehicleID;references:ID"`
	Mechanic_ID Mechanics `gorm:"foreignkey:MechanicID;references:ID"`
}

type Database struct {
	DB *gorm.DB
}
