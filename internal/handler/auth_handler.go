package handler

import (
	"encoding/json"
	apierrors "golang-for-devops/internal/errors"
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

// Login godoc
//
// @Summary Login
// @Description Authenticate a user and return JWT tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	var request models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	if request.Username == "" || request.Password == "" {
		// http.Error(w, "username and password are required", http.StatusBadRequest)
		apierrors.WriteError(
			w,
			apierrors.BadRequest("username and password are required"),
		)
		return
	}

	response, err := h.service.Login(
		request.Username,
		request.Password,
	)

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}
	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		response,
	)
}

func (h *AuthHandler) Refresh(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request models.RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	response, err := h.service.Refresh(
		request.RefreshToken,
	)

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	apierrors.WriteJSON(
		w,
		http.StatusCreated,
		response,
	)
}

func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request models.LogoutRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	err := h.service.Logout(
		request.RefreshToken,
	)

	if err != nil {
		apierrors.WriteError(
			w,
			err,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
