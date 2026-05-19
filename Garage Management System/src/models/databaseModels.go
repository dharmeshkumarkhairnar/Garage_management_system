package models

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	ID        uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Phone     uint64    `gorm:"column:phoneNumber" json:"phoneNumber"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type Vehicles struct {
	ID          uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	CustomerID  uint64    `gorm:"column:customer_id" json:"customer_id"`
	NumberPlate string    `gorm:"column:number_plate;uniqueIndex" json:"number_plate"`
	Model       string    `gorm:"column:model" json:"model"`
	Created_at  time.Time `gorm:"column:created_at" json:"created_at"`

	CustId Customers `gorm:"foreignkey:CustomerID;references:ID"`
}

type Mechanics struct {
	ID           uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	AadharNumber string    `gorm:"column:aadhar_number;uniqueIndex" json:"aadhar_number"`
	Phone        uint64    `gorm:"column:phoneNumber" json:"phoneNumber"`
	Created_at   time.Time `gorm:"column:created_at" json:"created_at"`
}

type ServiceMaster struct {
	ID      uint64  `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Service string  `gorm:"column:service; not null" json:"service"`
	Amount  float64 `gorm:"column:amount" json:"amount"`
}

type VisitRecords struct {
	ID           uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	VehicleID    uint64    `gorm:"column:vehicle_id" json:"vehicle_id"`
	MechanicID   uint64    `gorm:"column:mechanic_id" json:"mechanic_id"`
	ArrivalDate  time.Time `gorm:"column:arrival_date;type:date" json:"arrival_date"`
	DeliveryDate time.Time `gorm:"column:delivery_date;type:date" json:"delivery_date"`

	Vehicle_ID  Vehicles  `gorm:"foreignkey:VehicleID;references:ID"`
	Mechanic_ID Mechanics `gorm:"foreignkey:MechanicID;references:ID"`
}

type VisitServices struct {
	ID              uint64 `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	VisitRecordID   uint64 `gorm:"column:visit_record_id;uniqueIndex:idx_unique_rec" json:"visit_record_id"`
	ServiceMasterID uint64 `gorm:"column:service_master_id;uniqueIndex:idx_unique_rec" json:"service_master_id"`

	VisitRecord_ID   VisitRecords  `gorm:"foreignkey:VisitRecordID;references:ID"`
	ServiceMaster_ID ServiceMaster `gorm:"foreignkey:ServiceMasterID;references:ID"`
}

type Database struct {
	DB *gorm.DB
}
