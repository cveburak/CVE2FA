package service

import (
	"2fa/internal/model"
	"2fa/internal/repository"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type AccountService struct {
	accountRepo *repository.AccountRepo
}

func NewAccountService(accountRepo *repository.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (s *AccountService) List(userID int64) ([]model.Account, error) {
	accounts, err := s.accountRepo.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("list accounts: %w", err)
	}
	if accounts == nil {
		return []model.Account{}, nil
	}
	return accounts, nil
}

func (s *AccountService) Create(userID int64, req model.CreateAccountRequest) (*model.Account, error) {
	req.Issuer = strings.TrimSpace(req.Issuer)
	req.AccountName = strings.TrimSpace(req.AccountName)
	req.SecretBase32 = strings.TrimSpace(req.SecretBase32)

	if req.Issuer == "" {
		return nil, fmt.Errorf("issuer is required")
	}
	if req.AccountName == "" {
		return nil, fmt.Errorf("account name is required")
	}
	if req.SecretBase32 == "" {
		return nil, fmt.Errorf("secret key is required")
	}

	re := regexp.MustCompile(`^[A-Za-z2-7]+=*$`)
	if !re.MatchString(req.SecretBase32) {
		return nil, fmt.Errorf("invalid base32 secret key")
	}

	account := &model.Account{
		UserID:       userID,
		Issuer:       req.Issuer,
		AccountName:  req.AccountName,
		SecretBase32: req.SecretBase32,
		CreatedAt:    time.Now(),
	}

	return s.accountRepo.Create(account)
}

func (s *AccountService) Delete(id, userID int64) error {
	return s.accountRepo.Delete(id, userID)
}

func (s *AccountService) GetSecret(id, userID int64) (string, error) {
	account, err := s.accountRepo.FindByID(id, userID)
	if err != nil {
		return "", fmt.Errorf("get secret: %w", err)
	}
	if account == nil {
		return "", fmt.Errorf("account not found")
	}
	return account.SecretBase32, nil
}
