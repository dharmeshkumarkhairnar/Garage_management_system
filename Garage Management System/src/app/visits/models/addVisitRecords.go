package models

type BFFAddVisitRecordsRequest struct {
	NumberPlate  string   `json:"number_plate" example:"MH19BW3626" validate:"required,numPlateFormat"`
	MechanicId   string   `json:"mechanic_id" example:"firstname012" validate:"required"`
	ArrivalDate  string   `json:"arrival_date" example:"2026-05-20"`
	DeliveryDate string   `json:"delivery_date" example:"2026-05-20"`
	Services     []string `json:"services" example:"[painting,denting,oiling]"`
}

type BFFAddVisitRecordsResponse struct {
	Status string `json:"status" example:"Visit Records Added Sucessfully"`
}
