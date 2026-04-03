package handlers

import (
	"encoding/json"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/service"
	"net/http"
)

type PinHandler struct {
	service *service.PinService
}

func NewPinHandler(s *service.PinService) *PinHandler {
	return &PinHandler{s}
}

func (h *PinHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pins, err := h.service.Index(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		json.NewEncoder(w).Encode(resp)
	}

	w.WriteHeader(http.StatusOK)
	resp := models.DataResponse[models.Pin]{
		Message: "Success",
		Data:    pins,
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *PinHandler) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pin := models.Pin{
		Description: r.FormValue("description"),
	}

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   "File is too big",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)

		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)

		return
	}
	defer file.Close()

	err = h.service.SavePin(r.Context(), &pin, r.Context().Value("user").(string), file, header)
	if err != nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)

		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := models.Response{
		Message: "Pin created",
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *PinHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   "Description or image required",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	var pin models.Pin
	err := json.NewDecoder(r.Body).Decode(&pin)

	err = h.service.Update(r.Context(), r.PathValue("id"), r.Context().Value("user").(string), &pin)

	if err != nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.Response{
		Message: "Pin updated",
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *PinHandler) Remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := h.service.Remove(r.Context(), r.PathValue("id"), r.Context().Value("user").(string))

	if err != nil {
		resp := models.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp := models.Response{
		Message: "Pin removed",
	}
	json.NewEncoder(w).Encode(resp)
}
