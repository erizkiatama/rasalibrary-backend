package auth

import (
	"time"

	"github.com/erizkiatama/rasalibrary-backend/auth/helper"
	"github.com/erizkiatama/rasalibrary-backend/models"
)

var PARSE_DATE_FORMAT = "2006-01-02"

type service struct {
	repo models.AuthRepository
}

func NewAuthService(repo models.AuthRepository) models.AuthService {
	return &service{
		repo: repo,
	}
}

func (ths *service) Login(req models.LoginRequest) (*models.TokenPair, error) {
	user, err := ths.repo.GetUserWithProfileByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = helper.ComparePassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(user.ID, user.Profile.ID)
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

	dob, err := time.Parse(PARSE_DATE_FORMAT, req.DateOfBirth)
	if err != nil {
		return nil, models.NewClientError(
			"0100014",
			"wrong format for date of birth",
			400,
		)
	}

	newUser := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		IsAdmin:  req.IsAdmin,
		Profile: models.UserProfile{
			Name:         req.Name,
			DateOfBirth:  dob,
			Address:      req.Address,
			Sex:          req.Sex,
			PhoneNumber:  req.PhoneNumber,
			ProfilePhoto: req.ProfilePhoto,
		},
	}

	userID, profileID, err := ths.repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	tokenPair, err := helper.GenerateTokenPair(userID, profileID)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
