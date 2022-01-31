package models

import (
	"time"
)

type User struct {
	ID        uint      `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsAdmin   bool      `db:"is_admin"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AuthService interface {
	Login(LoginRequest) (*TokenPair, error)
	Register(RegisterRequest) (*TokenPair, error)
}

type AuthRepository interface {
	CreateUser(*User) error
	GetUserByEmail(string) (*User, error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type TokenPair struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}
