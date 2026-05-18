package models

type BFFCreateCustomerRequest struct {
	Name            string `json:"customerName" example:"Arjit" validate:"required,min=5,max=32"`
	Email           string `json:"email" example:"arijit@gmail.com" validate:"required,email"`
	PhoneNumber     string `json:"phoneNumber" example:"7568912340" validate:"required,len=10,numeric"`
	Password        string `json:"password" example:"Admin@123" validate:"required,min=8,strongPassword,max=20"`
	ConfirmPassword string `json:"confirmPassword" gorm:"column:confirmPassword" example:"Admin@123" validate:"required,min=8,eqfield=Password"`
}

type BFFCreateCustomerResponse struct {
	Message string `json:"message" example:"customer created successfully"`
}

type ErrorMessage struct {
	Key          string `json:"key,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

type ErrorAPIResponse struct {
	Message ErrorMessage `json:"errors,omitempty"`
	Error   string       `json:"error,omitempty"`
}
