package request

import "myapp-backend/service/model"

type AddUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"omitempty"`
}

func (addUserRequest AddUserRequest) ToModel() model.UserCreate {
	return model.UserCreate{
		Username: addUserRequest.Username,
		Email:    addUserRequest.Email,
		Password: addUserRequest.Password,
		Role:     addUserRequest.Role,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateBalanceRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type TransactionRequest struct {
	FromUserID int64   `json:"from_user_id"`
	ToUserID   int64   `json:"to_user_id"`
	Amount     float64 `json:"amount"`
}

type Movie struct {
	Title    string
	Director string
	Year     int
	Genre    string
}
