package models

type BFFCreateVehicleRequest struct {
	NumberPlate string `json:"number_plate" example:"MH19BW3626" validate:"required,numPlateFormat"`
	Model       string `json:"model" example:"Dezire" validate:"required"`
}

type BFFCreateVehicleResponse struct {
	Status string `json:"status" example:"Successfully Created"`
}
