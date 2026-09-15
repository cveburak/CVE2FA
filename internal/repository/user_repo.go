package repository

import (
	"2fa/internal/model"
	"database/sql"
	"fmt"
	"time"
)

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	row := DB.QueryRow(
		`SELECT id, username, password_hash, must_change_password, created_at FROM users WHERE username = ?`,
		username,
	)

	u := &model.User{}
	var mustChange int
	err := row.Scan(&u.ID, &u.Username, &u.PasswordHash, &mustChange, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}
	u.MustChangePassword = mustChange == 1
	return u, nil
}

func (r *UserRepo) FindByID(id int64) (*model.User, error) {
	row := DB.QueryRow(
		`SELECT id, username, password_hash, must_change_password, created_at FROM users WHERE id = ?`,
		id,
	)

	u := &model.User{}
	var mustChange int
	err := row.Scan(&u.ID, &u.Username, &u.PasswordHash, &mustChange, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("find user by id: %w", err)
	}
	u.MustChangePassword = mustChange == 1
	return u, nil
}

func (r *UserRepo) Create(username, passwordHash string) (*model.User, error) {
	now := time.Now()
	res, err := DB.Exec(
		`INSERT INTO users (username, password_hash, must_change_password, created_at) VALUES (?, ?, 0, ?)`,
		username, passwordHash, now,
	)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	id, _ := res.LastInsertId()
	return &model.User{
		ID:                 id,
		Username:           username,
		PasswordHash:       passwordHash,
		MustChangePassword: false,
		CreatedAt:          now,
	}, nil
}

func (r *UserRepo) UpdatePassword(id int64, hash string) error {
	_, err := DB.Exec(
		`UPDATE users SET password_hash = ?, must_change_password = 0 WHERE id = ?`,
		hash, id,
	)
	return err
}

func (r *UserRepo) SeedAdmin() error {
	existing, _ := r.FindByUsername("admin")
	if existing != nil {
		return nil
	}

	hash := "$2a$10$dummyhashforadminuserpleasechange"
	now := time.Now()
	_, err := DB.Exec(
		`INSERT INTO users (username, password_hash, must_change_password, created_at) VALUES (?, ?, 1, ?)`,
		"admin", hash, now,
	)
	if err != nil {
		return fmt.Errorf("seed admin: %w", err)
	}
	return nil
}

func (r *UserRepo) SetMustChangePassword(id int64, val bool) error {
	v := 0
	if val {
		v = 1
	}
	_, err := DB.Exec(`UPDATE users SET must_change_password = ? WHERE id = ?`, v, id)
	return err
}
