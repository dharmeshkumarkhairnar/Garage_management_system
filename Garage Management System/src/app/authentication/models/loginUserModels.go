package models

type BFFLoginUserRequest struct {
	Email    string `json:"email" exampple:"arijit@gmail.com" validate:"required"`
	Password string `json:"password" example:"Admin@123" validate:"required"`
	Role     string `json:"role" example:"customer or admin" validate:"required"`
}

type BFFLoginUserResponse struct {
	Token   string `json:"token"`
	Message string `json:"message" example:"user logged in successfully"`
}
