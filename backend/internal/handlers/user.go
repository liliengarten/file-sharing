package handlers

import (
	"net/http"
	"encoding/json"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/validator"
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

type ErrorResponse struct {
	Message string `json:"message"`
	Error string `json:"error"`
}



func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validationErr := validator.Validate(user)

	//Ошибка валидации
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(validationErr)
		return
	}

	err = h.service.Register(r.Context(), user)
	
	//Ошибка сервиса
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		
		resp := ErrorResponse{
			Message: "Error",
			Error: err.Error(),
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	//Успех
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
