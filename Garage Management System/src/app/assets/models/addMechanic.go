package models

type BFFAddMechanicRequest struct {
	Name         string `json:"name" example:"Mukesh roy"`
	AddharNumber string `json:"aadhar_number" example:"112233445566"`
	Phone        string `json:"phone" example:"9881463919"`
}

type BFFAddMechanicResponse struct {
	Status string `json:"status" example:"Successfully added"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"errors,omitempty"`
	Error   string       `json:"error,omitempty"`
}
