package handlers

import (
	"encoding/json"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/responder"
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
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validationErr := validator.Validate(user)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(validationErr)
		return
	}

	err = h.service.Register(r.Context(), user)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Register successfully", http.StatusCreated)
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
		responder.ErrorResponse(w, "Authentification failed", http.StatusBadRequest)
		return
	}

	responder.Response(w, token, http.StatusOK)
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	profile, err := h.service.GetProfile(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.DataResponse(w, "Success", profile, http.StatusOK)
}
