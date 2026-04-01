package models

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
type DataResponse[T any] struct {
	Message string `json:"message"`
	Data    []T    `json:"data"`
}
