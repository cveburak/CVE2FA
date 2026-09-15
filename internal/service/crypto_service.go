package service

import (
	"2fa/internal/model"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

const (
	SaltLen      = 16
	NonceLen     = 12
	KeyLen       = 32
	ArgonTime    = 3
	ArgonMem     = 64 * 1024
	ArgonThreads = 4
)

type EncryptedPayload struct {
	Version   int    `json:"version"`
	SaltB64   string `json:"salt"`
	NonceB64  string `json:"nonce"`
	CipherB64 string `json:"ciphertext"`
}

type CryptoService struct{}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

func (s *CryptoService) Encrypt(plaintext []byte, masterPassword string) ([]byte, error) {
	salt := make([]byte, SaltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("generate salt: %w", err)
	}

	key := argon2.IDKey([]byte(masterPassword), salt, ArgonTime, ArgonMem, ArgonThreads, KeyLen)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("new cipher: %w", err)
	}

	nonce := make([]byte, NonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	payload := EncryptedPayload{
		Version:   1,
		SaltB64:   base64.StdEncoding.EncodeToString(salt),
		NonceB64:  base64.StdEncoding.EncodeToString(nonce),
		CipherB64: base64.StdEncoding.EncodeToString(ciphertext),
	}

	return json.Marshal(payload)
}

func (s *CryptoService) Decrypt(encrypted []byte, masterPassword string) ([]byte, error) {
	var payload EncryptedPayload
	if err := json.Unmarshal(encrypted, &payload); err != nil {
		return nil, fmt.Errorf("parse encrypted payload: %w", err)
	}

	salt, err := base64.StdEncoding.DecodeString(payload.SaltB64)
	if err != nil {
		return nil, fmt.Errorf("decode salt: %w", err)
	}

	nonce, err := base64.StdEncoding.DecodeString(payload.NonceB64)
	if err != nil {
		return nil, fmt.Errorf("decode nonce: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(payload.CipherB64)
	if err != nil {
		return nil, fmt.Errorf("decode ciphertext: %w", err)
	}

	key := argon2.IDKey([]byte(masterPassword), salt, ArgonTime, ArgonMem, ArgonThreads, KeyLen)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("new cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("new gcm: %w", err)
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decrypt failed (wrong password or corrupted data): %w", err)
	}

	return plaintext, nil
}

func (s *CryptoService) ExportData(accounts []model.Account, masterPassword string) ([]byte, error) {
	export := model.ExportData{
		Version:  1,
		Accounts: make([]model.ExportAccount, len(accounts)),
	}

	for i, a := range accounts {
		export.Accounts[i] = model.ExportAccount{
			Issuer:       a.Issuer,
			AccountName:  a.AccountName,
			SecretBase32: a.SecretBase32,
		}
	}

	plaintext, err := json.Marshal(export)
	if err != nil {
		return nil, fmt.Errorf("marshal export: %w", err)
	}

	return s.Encrypt(plaintext, masterPassword)
}

func (s *CryptoService) ImportData(encrypted []byte, masterPassword string) (*model.ExportData, error) {
	plaintext, err := s.Decrypt(encrypted, masterPassword)
	if err != nil {
		return nil, fmt.Errorf("import decrypt: %w", err)
	}

	var data model.ExportData
	if err := json.Unmarshal(plaintext, &data); err != nil {
		return nil, fmt.Errorf("parse import data: %w", err)
	}

	return &data, nil
}
