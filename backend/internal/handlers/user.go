package handlers

import (
	"net/http"
	"encoding/json"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/models"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}



type UserResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	User UserResponse `json:"user"`
}



func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	resp := RegisterResponse{
		Message: "Success",
		User: UserResponse {
			Username: user.Username,
			Email: user.Email,
		},
	}

	json.NewEncoder(w).Encode(resp)
}
