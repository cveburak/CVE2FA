package model

import "time"

type Account struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"userId"`
	Issuer       string    `json:"issuer"`
	AccountName  string    `json:"accountName"`
	SecretBase32 string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}

type CreateAccountRequest struct {
	Issuer       string `json:"issuer"`
	AccountName  string `json:"accountName"`
	SecretBase32 string `json:"secretBase32"`
}

type CodeResponse struct {
	Code      string `json:"code"`
	Remaining int    `json:"remaining"`
	Period    int    `json:"period"`
}
