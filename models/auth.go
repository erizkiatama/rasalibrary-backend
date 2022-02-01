package models

import (
	"time"
)

type User struct {
	ID        uint
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Profile   UserProfile
}

type UserProfile struct {
	ID           uint
	Name         string
	DateOfBirth  time.Time
	Address      string
	Sex          string
	PhoneNumber  string
	ProfilePhoto string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AuthService interface {
	Login(LoginRequest) (*TokenPair, error)
	Register(RegisterRequest) (*TokenPair, error)
}

type AuthRepository interface {
	CreateUser(User) (uint, uint, error)
	GetUserWithProfileByEmail(string) (*User, error)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsAdmin      bool   `json:"is_admin"`
	Name         string `json:"name"`
	DateOfBirth  string `json:"dob"`
	Address      string `json:"address"`
	Sex          string `json:"sex"`
	PhoneNumber  string `json:"phone_number"`
	ProfilePhoto string `json:"profile_photo"`
}

type TokenPair struct {
	Access  string `json:"access,omitempty"`
	Refresh string `json:"refresh,omitempty"`
}
