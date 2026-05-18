package models

type BFFCreateVehicleRequest struct {
	CustomerID  string    `json:"customer_id"`
	NumberPlate string    `json:"number_plate"`
	Model       string    `json:"model"`
}

type BFFCreateVehicleResponse struct {
	Status string `json:"status"`
}

