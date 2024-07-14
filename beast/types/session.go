package types

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Token     string `json:"token"`
	UserID    string `json:"user_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	ExpiresAt string `json:"expires_at"`
}
