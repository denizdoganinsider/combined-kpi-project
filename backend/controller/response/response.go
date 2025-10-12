package response

import "myapp-backend/domain"

type ErrorResponse struct {
	Message string `json:"error_description"`
}

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type UserResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func ToResponse(user domain.User) UserResponse {
	return UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ToResponseList(allUsers []domain.User) []UserResponse {
	var userResponseList = []UserResponse{}
	for _, user := range allUsers {
		userResponseList = append(userResponseList, ToResponse(user))
	}

	return userResponseList
}
