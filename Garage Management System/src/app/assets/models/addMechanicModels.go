package models

type BFFAddMechanicRequest struct {
	Name         string `json:"name" example:"Mukesh Roy" validate:"required,alphaspace"`
	AadharNumber string `json:"aadhar_number" example:"112233445566" validate:"required,aadharformat"`
	PhoneNumber  uint64 `json:"phoneNumber" gorm:"column:phoneNumber" example:"8432805566" validate:"required,min=1000000000,max=9999999999"`
}

type BFFAddMechanicResponse struct {
	Status string `json:"status" example:"Successfully added"`
}
