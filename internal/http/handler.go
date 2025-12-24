package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alexaaaant/user-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(us *service.UserService) *Handler {
	return &Handler{userService: us}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email string `json:"email"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
