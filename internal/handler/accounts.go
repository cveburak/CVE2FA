package handler

import (
	"2fa/internal/model"
	"2fa/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type AccountHandler struct {
	accountService *service.AccountService
	totpService    *service.TOTPService
}

func NewAccountHandler(accountService *service.AccountService, totpService *service.TOTPService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
		totpService:    totpService,
	}
}

func (h *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	accounts, err := h.accountService.List(userID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, accounts)
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)

	var req model.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	req.SecretBase32 = strings.ToUpper(strings.TrimSpace(req.SecretBase32))

	account, err := h.accountService.Create(userID, req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, account)
}

func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	if err := h.accountService.Delete(id, userID); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "deleted"})
}

func (h *AccountHandler) GetCode(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int64)
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid account id")
		return
	}

	secret, err := h.accountService.GetSecret(id, userID)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	code, err := h.totpService.GenerateCode(secret)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, code)
}
