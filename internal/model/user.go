package model

import "time"

type User struct {
	ID                 int64     `json:"id"`
	Username           string    `json:"username"`
	PasswordHash       string    `json:"-"`
	MustChangePassword bool      `json:"mustChangePassword"`
	CreatedAt          time.Time `json:"createdAt"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token              string `json:"token,omitempty"`
	MustChangePassword bool   `json:"mustChangePassword"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
