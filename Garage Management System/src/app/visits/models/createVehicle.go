package models

type BFFCreateVehicleRequest struct {
	CustomerID  uint64    `json:"customer_id" example:"12"`
	NumberPlate string    `json:"number_plate" example:"MH19BW3626"`
	Model       string    `json:"model" example:"Dezire"`
}

type BFFCreateVehicleResponse struct {
	Status string `json:"status" example:"Successfully created"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"errors,omitempty"`
	Error   string       `json:"error,omitempty"`
}