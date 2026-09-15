package service

import (
	"2fa/internal/model"
	"2fa/internal/totp"
	"encoding/base32"
	"fmt"
	"strings"
	"time"
)

type TOTPService struct{}

func NewTOTPService() *TOTPService {
	return &TOTPService{}
}

func (s *TOTPService) GenerateCode(secretBase32 string) (*model.CodeResponse, error) {
	secret, err := decodeSecret(secretBase32)
	if err != nil {
		return nil, fmt.Errorf("decode secret: %w", err)
	}

	now := time.Now()
	code := totp.GenerateCodeFromTime(secret, now)
	remaining := totp.GetRemainingSeconds(now)

	return &model.CodeResponse{
		Code:      code,
		Remaining: remaining,
		Period:    totp.Period,
	}, nil
}

func (s *TOTPService) ValidateCode(secretBase32, code string) bool {
	secret, err := decodeSecret(secretBase32)
	if err != nil {
		return false
	}
	return totp.ValidateCode(secret, code, time.Now())
}

func decodeSecret(s string) ([]byte, error) {
	s = strings.ToUpper(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimRight(s, "=")

	decoded, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(s)
	if err != nil {
		decoded, err = base32.StdEncoding.DecodeString(s)
		if err != nil {
			return nil, fmt.Errorf("base32 decode: %w", err)
		}
	}
	return decoded, nil
}
