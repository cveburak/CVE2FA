package handler

import (
	"2fa/internal/service"
	"encoding/json"
	"net/http"
)

type ExportHandler struct {
	accountService *service.AccountService
	cryptoService  *service.CryptoService
}

func NewExportHandler(accountService *service.AccountService, cryptoService *service.CryptoService) *ExportHandler {
	return &ExportHandler{
		accountService: accountService,
		cryptoService:  cryptoService,
	}
}

func (h *ExportHandler) Export(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	var req struct {
		MasterPassword string `json:"masterPassword"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if len(req.MasterPassword) < 4 {
		writeError(w, http.StatusBadRequest, "master password must be at least 4 characters")
		return
	}

	accounts, err := h.accountService.List(userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	encrypted, err := h.cryptoService.ExportData(accounts, req.MasterPassword)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "encryption failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=CVE2fa.2fadata")
	w.WriteHeader(http.StatusOK)
	w.Write(encrypted)
}
