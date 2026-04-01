package handlers

import (
	"encoding/json"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/validator"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
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

		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	//Успех
	w.WriteHeader(http.StatusCreated)

	resp := models.RegisterResponse{
		Message: "Success",
		User: models.UserResponse{
			Username: user.Username,
			Email:    user.Email,
		},
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), user)

	if err != nil {
		resp := models.Response{
			Message: "Authentification failed",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)

		return
	}

	resp := models.LoginResponse{
		Message: "Authentification succeed",
		Token:   token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
