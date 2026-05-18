package models

type BFFCreateCustomerRequest struct {
	ID              uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string `json:"customerName" example:"Arjit" validate:"required,min=5,max=32"`
	Email           string `json:"email" example:"arijit@gmail.com" validate:"required,email"`
	PhoneNumber     string `json:"phoneNumber" example:"7568912340" validate:"required,len=10,numeric"`
	Password        string `json:"password" example:"Secure@123" validate:"required,min=8,strongPassword,max=20"`
	ConfirmPassword string `json:"confirmPassword" gorm:"column:confirmPassword" example:"Secure@123" validate:"required,min=8,eqfield=Password"`
}

type BFFCreateCustomerResponse struct {
	Message string `json:"message" example:"customer created successfully"`
}
