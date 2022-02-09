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

	sex := "M"
	if req.Sex != "Male" {
		sex = "F"
	}

	newUser := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		IsAdmin:  req.IsAdmin,
		Profile: models.UserProfile{
			Name:         req.Name,
			DateOfBirth:  dob,
			Address:      req.Address,
			Sex:          sex,
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

func (ths *service) GetLoggedInUser(userID uint) (*models.GetUserResponse, error) {
	user, err := ths.repo.GetUserWithProfileByID(userID)
	if err != nil {
		return nil, err
	}

	sex := "Male"
	if user.Profile.Sex != "M" {
		sex = "Female"
	}

	response := &models.GetUserResponse{
		Email:        user.Email,
		IsAdmin:      user.IsAdmin,
		Name:         user.Profile.Name,
		DateOfBirth:  user.Profile.DateOfBirth.Format(PARSE_DATE_FORMAT),
		Address:      user.Profile.Address,
		Sex:          sex,
		PhoneNumber:  user.Profile.PhoneNumber,
		ProfilePhoto: user.Profile.ProfilePhoto,
	}

	return response, nil
}
