package auth

import (
	"github.com/erizkiatama/rasalibrary-backend/auth/helper"
	"github.com/erizkiatama/rasalibrary-backend/models"
)

type service struct {
	repo models.AuthRepository
}

func NewAuthService(repo models.AuthRepository) models.AuthService {
	return &service{
		repo: repo,
	}
}

func (ths *service) Login(req models.LoginRequest) (*models.TokenPair, error) {
	user, err := ths.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (ths *service) Register(req models.RegisterRequest) (*models.TokenPair, error) {
	hashedPassword, err := helper.EncryptPassword([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
		IsAdmin:  req.IsAdmin,
	}

	err = ths.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(newUser.ID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
