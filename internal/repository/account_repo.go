package repository

import (
	"2fa/internal/model"
	"database/sql"
	"fmt"
)

type AccountRepo struct{}

func NewAccountRepo() *AccountRepo {
	return &AccountRepo{}
}

func (r *AccountRepo) FindByUserID(userID int64) ([]model.Account, error) {
	rows, err := DB.Query(
		`SELECT id, user_id, issuer, account_name, secret_base32, created_at FROM accounts WHERE user_id = ? ORDER BY issuer ASC`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("find accounts: %w", err)
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var a model.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.Issuer, &a.AccountName, &a.SecretBase32, &a.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan account: %w", err)
		}
		accounts = append(accounts, a)
	}
	return accounts, rows.Err()
}

func (r *AccountRepo) FindByID(id, userID int64) (*model.Account, error) {
	row := DB.QueryRow(
		`SELECT id, user_id, issuer, account_name, secret_base32, created_at FROM accounts WHERE id = ? AND user_id = ?`,
		id, userID,
	)

	var a model.Account
	err := row.Scan(&a.ID, &a.UserID, &a.Issuer, &a.AccountName, &a.SecretBase32, &a.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find account: %w", err)
	}
	return &a, nil
}

func (r *AccountRepo) Create(a *model.Account) (*model.Account, error) {
	res, err := DB.Exec(
		`INSERT INTO accounts (user_id, issuer, account_name, secret_base32, created_at) VALUES (?, ?, ?, ?, ?)`,
		a.UserID, a.Issuer, a.AccountName, a.SecretBase32, a.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	id, _ := res.LastInsertId()
	a.ID = id
	return a, nil
}

func (r *AccountRepo) Delete(id, userID int64) error {
	res, err := DB.Exec(`DELETE FROM accounts WHERE id = ? AND user_id = ?`, id, userID)
	if err != nil {
		return fmt.Errorf("delete account: %w", err)
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return fmt.Errorf("account not found")
	}
	return nil
}

func (r *AccountRepo) FindAll() ([]model.Account, error) {
	rows, err := DB.Query(
		`SELECT id, user_id, issuer, account_name, secret_base32, created_at FROM accounts ORDER BY issuer ASC`,
	)
	if err != nil {
		return nil, fmt.Errorf("find all accounts: %w", err)
	}
	defer rows.Close()

	var accounts []model.Account
	for rows.Next() {
		var a model.Account
		if err := rows.Scan(&a.ID, &a.UserID, &a.Issuer, &a.AccountName, &a.SecretBase32, &a.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan account: %w", err)
		}
		accounts = append(accounts, a)
	}
	return accounts, rows.Err()
}
