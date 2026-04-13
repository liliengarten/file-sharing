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
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validationErr := validator.Validate(user)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)

		err = json.NewEncoder(w).Encode(validationErr)
		if err != nil {
			responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

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
	var user models.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), user)
	if err != nil {
		responder.ErrorResponse(w, "authentification failed", http.StatusBadRequest)
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

func (h *UserHandler) GetLikes(w http.ResponseWriter, r *http.Request) {
	pins, err := h.service.GetLikes(r.Context(), r.URL.Query().Get("page"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.DataResponse(w, "Success", pins, http.StatusOK)
}
