package models

type BFFDeleteServiceRequest struct {
	Service string `json:"service" example:"painting" validate:"required,min=3"`
}

type BFFDeleteerviceResponse struct {
	Message string `json:"message" example:"service deleted successfully"`
}
