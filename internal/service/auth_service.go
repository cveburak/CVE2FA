package service

import (
	"2fa/internal/model"
	"2fa/internal/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repository.UserRepo
	jwtSecret []byte
}

func NewAuthService(userRepo *repository.UserRepo) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: []byte("change-me-in-production-CVE2fa-2026"),
	}
}

func (s *AuthService) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	token, err := s.generateToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}

	return &model.LoginResponse{
		Token:              token,
		MustChangePassword: user.MustChangePassword,
	}, nil
}

func (s *AuthService) ChangePassword(userID int64, req model.ChangePasswordRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.OldPassword)); err != nil {
		return fmt.Errorf("current password is incorrect")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	return s.userRepo.UpdatePassword(userID, string(hash))
}

func (s *AuthService) SeedAdmin() error {
	hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash admin password: %w", err)
	}

	existing, _ := s.userRepo.FindByUsername("admin")
	if existing != nil {
		return nil
	}

	now := time.Now()
	_, err = repository.DB.Exec(
		`INSERT INTO users (username, password_hash, must_change_password, created_at) VALUES (?, ?, 1, ?)`,
		"admin", string(hash), now,
	)
	if err != nil {
		return fmt.Errorf("seed admin: %w", err)
	}
	return nil
}

func (s *AuthService) generateToken(userID int64, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *AuthService) ValidateToken(tokenStr string) (int64, string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	sub := int64(claims["sub"].(float64))
	username := claims["username"].(string)
	return sub, username, nil
}
