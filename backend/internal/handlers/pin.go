package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/responder"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/validator"
)

type PinHandler struct {
	service *service.PinService
}

func NewPinHandler(s *service.PinService) *PinHandler {
	return &PinHandler{s}
}

func (h *PinHandler) Index(w http.ResponseWriter, r *http.Request) {
	pins, err := h.service.Index(r.Context(), r.URL.Query().Get("page"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	responder.DataResponse(w, "Success", pins, http.StatusOK)
}

func (h *PinHandler) UserPins(w http.ResponseWriter, r *http.Request) {
	pins, err := h.service.UserPins(r.Context(), r.URL.Query().Get("page"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	responder.DataResponse(w, "Success", pins, http.StatusOK)
}

func (h *PinHandler) GetPin(w http.ResponseWriter, r *http.Request) {
	pin, err := h.service.GetPin(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.DataResponse(w, "Success", pin, http.StatusOK)
}

func (h *PinHandler) Add(w http.ResponseWriter, r *http.Request) {
	var pin models.Pin
	pin.Description = r.PostFormValue("description")

	validationErr := validator.Validate(pin)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		responder.ErrorResponse(w, "File is too big", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = h.service.SavePin(r.Context(), &pin, r.Context().Value("user").(string), file, header)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin created", http.StatusCreated)
}

func (h *PinHandler) Update(w http.ResponseWriter, r *http.Request) {
	var pin models.Pin
	err := json.NewDecoder(r.Body).Decode(&pin)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	pinID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	pin.ID = pinID

	validationErr := validator.Validate(pin)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	err = h.service.Update(r.Context(), &pin)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin updated", http.StatusOK)
}

func (h *PinHandler) Remove(w http.ResponseWriter, r *http.Request) {
	err := h.service.Remove(r.Context(), r.PathValue("id"), r.Context().Value("user").(string))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin removed", http.StatusOK)
}

func (h *PinHandler) LikePin(w http.ResponseWriter, r *http.Request) {
	err := h.service.LikePin(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin liked", http.StatusOK)
}

func (h *PinHandler) UnlikePin(w http.ResponseWriter, r *http.Request) {
	err := h.service.UnlikePin(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin unliked", http.StatusOK)
}
