package models

import "time"

type BFFAddVisitRecordsRequest struct {
	NumberPlate  string    `json:"number_plate" example:"MH19BW3626" validate:"required,numPlateFormat"`
	MechanicId   string    `json:"mechanic_id" example:"firstname012" validate:"required,mechIDFormat"`
	ArrivalDate  time.Time `json:"arrival_date" example:"2026-05-20"`
	DeliveryDate time.Time `json:"delivery_date" example:"2026-05-20"`
	Services     []string  `json:"services" example:"[painting,denting,oiling]"`
}

type BFFVisitRecordsResponse struct {
	Status string `json:"status" example:"Visit Records Added Sucessfully"`
}
