package handler

import (
	"encoding/json"
	"golang-for-devops/internal/models"
	"golang-for-devops/internal/service"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Login end point
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var request models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if request.Username == "" || request.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(
		request.Username,
		request.Password,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := models.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
