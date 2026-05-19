package models

type BFFAddMechanicRequest struct {
	Name         string `json:"name" example:"Mukesh roy"`
	AddharNumber string `json:"aadhar_number" example:"112233445566"`
	Phone        uint64 `json:"phone" example:"9881463919"`
}

type BFFAddMechanicResponse struct {
	Status string `json:"status" example:"Successfully added"`
}
