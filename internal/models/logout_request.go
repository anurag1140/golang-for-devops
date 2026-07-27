package models

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}
