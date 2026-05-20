package models

type BFFCreateVehicleRequest struct {
	NumberPlate string `json:"number_plate" example:"MH19BW3626" validate:"required,numPlateFormat"`
	Model       string `json:"model" example:"Dezire" validate:"required"`
}

type BFFCreateVehicleResponse struct {
	Status string `json:"status" example:"Successfully Created"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"errors,omitempty"`
	Error   string       `json:"error,omitempty"`
}
