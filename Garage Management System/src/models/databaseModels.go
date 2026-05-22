package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID        uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Email     string    `gorm:"column:email;uniqueIndex" json:"email"`
	Phone     uint64    `gorm:"column:phoneNumber" json:"phoneNumber"`
	Role      string    `gorm:"column:role;default:customer" json:"role"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type Vehicles struct {
	ID          uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	UserID      uint64    `gorm:"column:user_id" json:"user_id"`
	NumberPlate string    `gorm:"column:number_plate;uniqueIndex" json:"number_plate"`
	Model       string    `gorm:"column:model" json:"model"`
	Created_at  time.Time `gorm:"column:created_at" json:"created_at"`

	UId Users `gorm:"foreignkey:UserID;references:ID"`
}

type Mechanics struct {
	MechID       string    `gorm:"column:mech_id;primarykey" json:"mech_id"`
	Name         string    `gorm:"column:name" json:"name"`
	AadharNumber string    `gorm:"column:aadhar_number;uniqueIndex" json:"aadhar_number"`
	Phone        uint64    `gorm:"column:phoneNumber" json:"phoneNumber"`
	Created_at   time.Time `gorm:"column:created_at" json:"created_at"`
}

type ServiceMaster struct {
	ID      uint64  `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	Service string  `gorm:"column:service; uniqueIndex; not null" json:"service"`
	Amount  float64 `gorm:"column:amount" json:"amount"`
}

type VisitRecords struct {
	ID           uint64    `gorm:"column:id;primarykey;autoIncrement" json:"id"`
	VehicleID    uint64    `gorm:"column:vehicle_id" json:"vehicle_id"`
	MechanicID   string    `gorm:"column:mech_id" json:"mech_id"`
	ArrivalDate  time.Time `gorm:"column:arrival_date;type:date" json:"arrival_date"`
	DeliveryDate time.Time `gorm:"column:delivery_date;type:date;check:delivery_date>=arrival_date" json:"delivery_date"`

	Vehicle_ID  Vehicles  `gorm:"foreignkey:VehicleID;references:ID"`
	Mechanic_ID Mechanics `gorm:"foreignkey:MechanicID;references:MechID"`
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
