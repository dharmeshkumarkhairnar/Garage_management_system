package models

type BFFCreateVehicleRequest struct {
	CustomerID  uint64    `json:"customer_id"`
	NumberPlate string    `json:"number_plate"`
	Model       string    `json:"model"`
}

type BFFCreateVehicleResponse struct {
	Status string `json:"status"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"errors,omitempty"`
	Error   string       `json:"error,omitempty"`
}