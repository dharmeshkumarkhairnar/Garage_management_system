package models

type BFFDeleteMechanicRequest struct {
	MechanicID string `json:"mechanic_id" example:"arijit709" validate:"required"`
}

type BFFDeleteMechanicResponse struct {
	Status string `json:"status" example:"Successfully deleted"`
}
