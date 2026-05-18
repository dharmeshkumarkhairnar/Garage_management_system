package model

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	ID        uint64    `gorm:"column:id;primarykey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type Vehicles struct {
	ID          uint64    `gorm:"column:id;primarykey" json:"id"`
	CustomerID  string    `gorm:"column:customer_id" json:"customer_id"`
	NumberPlate string    `gorm:"column:number_plate;uniqueIndex" json:"number_plate"`
	Model       string    `gorm:"column:model" json:"model"`
	Created_at  time.Time `gorm:"column:created_at" json:"created_at"`

	CustId Customers `gorm:"foreignkey:CustomerID;references:ID"`
}

type Mechanics struct {
	ID    uint64 `gorm:"column:id;primarykey" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Phone string `gorm:"column:phone" json:"phone"`
}

type ServiceMaster struct {
	ID      uint64  `gorm:"column:id;primarykey" json:"id"`
	Service string  `gorm:"column:service; not null" json:"service"`
	Amount  float64 `gorm:"column:amount" json:"amount"`
}

type VisitRecords struct {
	ID           uint64    `gorm:"column:id;primarykey" json:"id"`
	VehicleID    uint64    `gorm:"column:vehicle_id" json:"vehicle_id"`
	MechanicID   uint64    `gorm:"column:mechanic_id;uniqueIndex" json:"mechanic_id"`
	ArrivalDate  time.Time `gorm:"column:arrival_date;type:date" json:"arrival_date"`
	DeliveryDate time.Time `gorm:"column:delivery_date;type:date" json:"delivery_date"`

	Vehicle_ID  Vehicles  `gorm:"foreignkey:VehicleID;references:ID"`
	Mechanic_ID Mechanics `gorm:"foreignkey:MechanicID;references:ID"`
}

type VisitServices struct {
	ID              uint64 `gorm:"column:id;primarykey" json:"id"`
	VisitRecordID   uint64 `gorm:"column:visit_record_id" json:"visit_record_id"`
	ServiceMasterID uint64 `gorm:"column:service_master_id" json:"service_master_id"`

	VisitRecord_ID   VisitRecords  `gorm:"foreignkey:VisitRecordID;references:ID"`
	ServiceMaster_ID ServiceMaster `gorm:"foreignkey:ServiceMasterID;references:ID"`
}

type Database struct {
	DB *gorm.DB
}
