package handlers

import (
	"encoding/json"
	"liliengarten/filesharing/internal/models"
	"liliengarten/filesharing/internal/responder"
	"liliengarten/filesharing/internal/service"
	"liliengarten/filesharing/internal/validator"
	"net/http"
)

type BoardHandler struct {
	service *service.BoardService
}

func NewBoardHandler(s *service.BoardService) *BoardHandler {
	return &BoardHandler{service: s}
}

func (h *BoardHandler) Index(w http.ResponseWriter, r *http.Request) {
	boards, err := h.service.Index(r.Context())
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	responder.DataResponse(w, "Success", boards, http.StatusOK)
}

func (h *BoardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var board models.Board
	err := json.NewDecoder(r.Body).Decode(&board)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	validationErr := validator.Validate(board)
	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErr)
		return
	}

	err = h.service.Create(r.Context(), &board)
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Board created successfully", http.StatusCreated)
}

func (h *BoardHandler) Remove(w http.ResponseWriter, r *http.Request) {
	err := h.service.Remove(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Board removed successfully", http.StatusOK)
}

func (h *BoardHandler) GetPins(w http.ResponseWriter, r *http.Request) {
	pins, err := h.service.GetPins(r.Context(), r.PathValue("id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.DataResponse(w, "Success", pins, http.StatusOK)
}

func (h *BoardHandler) AddPin(w http.ResponseWriter, r *http.Request) {
	err := h.service.AddPin(r.Context(), r.PathValue("board_id"), r.PathValue("pin_id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin added successfully", http.StatusCreated)
}

func (h *BoardHandler) RemovePin(w http.ResponseWriter, r *http.Request) {
	err := h.service.RemovePin(r.Context(), r.PathValue("board_id"), r.PathValue("pin_id"))
	if err != nil {
		responder.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	responder.Response(w, "Pin removed successfully", http.StatusCreated)
}
