package request

type RegisterRequest struct {
	Email    string `json:"email" validate:"email" example:"test@email.com"`
	Password string `json:"password" validate:"min=6" example:"test123"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email" example:"test@email.com"`
	Password string `json:"password" validate:"min=6" example:"test123"`
}
