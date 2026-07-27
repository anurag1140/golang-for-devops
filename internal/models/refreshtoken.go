package models

import "time"

type RefreshToken struct {
	Token     string
	Username  string
	ExpiresAt time.Time
	IsRevoked bool
}
