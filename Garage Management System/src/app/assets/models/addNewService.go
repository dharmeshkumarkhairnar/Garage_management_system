package models

type BFFAddNewServiceRequest struct {
	Service string  `json:"service" example:"painting" validate:"required,min=3"`
	Amount  float64 `json:"amount" example:"1745.99" validate:"min=99"`
}

type BFFAddNewServiceResponse struct {
	Message string `json:"message" example:"service added successfully"`
}
