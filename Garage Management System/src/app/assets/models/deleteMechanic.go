package models

type BFFDeleteMechanicRequest struct {
	AddharNumber string `json:"aadhar_number" example:"112233445566"`
}

type BFFDeleteMechanicResponse struct {
	Status string `json:"status" example:"Successfully deleted"`
}

