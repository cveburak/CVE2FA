package handler

import (
	"2fa/internal/model"
	"2fa/internal/repository"
	"2fa/internal/service"
	"fmt"
	"net/http"
	"time"
)

type ImportHandler struct {
	cryptoService *service.CryptoService
	accountRepo   *repository.AccountRepo
}

func NewImportHandler(cryptoService *service.CryptoService, accountRepo *repository.AccountRepo) *ImportHandler {
	return &ImportHandler{
		cryptoService: cryptoService,
		accountRepo:   accountRepo,
	}
}

func (h *ImportHandler) Import(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		writeError(w, http.StatusBadRequest, "invalid form data")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "file is required")
		return
	}
	defer file.Close()

	masterPassword := r.FormValue("masterPassword")
	if len(masterPassword) < 4 {
		writeError(w, http.StatusBadRequest, "master password must be at least 4 characters")
		return
	}

	buf := make([]byte, 10<<20)
	n, err := file.Read(buf)
	if err != nil && n == 0 {
		writeError(w, http.StatusBadRequest, "failed to read file")
		return
	}

	data, err := h.cryptoService.ImportData(buf[:n], masterPassword)
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("import failed: %v", err))
		return
	}

	imported := 0
	skipped := 0
	for _, ea := range data.Accounts {
		existing, _ := h.accountRepo.FindByUserID(userID)
		duplicate := false
		for _, e := range existing {
			if e.Issuer == ea.Issuer && e.AccountName == ea.AccountName {
				duplicate = true
				break
			}
		}
		if duplicate {
			skipped++
			continue
		}

		account := &model.Account{
			UserID:       userID,
			Issuer:       ea.Issuer,
			AccountName:  ea.AccountName,
			SecretBase32: ea.SecretBase32,
			CreatedAt:    time.Now(),
		}
		if _, err := h.accountRepo.Create(account); err != nil {
			continue
		}
		imported++
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"message":  "import completed",
		"imported": imported,
		"skipped":  skipped,
	})
}
