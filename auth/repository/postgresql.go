package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/erizkiatama/rasalibrary-backend/models"
	"github.com/jmoiron/sqlx"
)

type postgreSQLRepository struct {
	db *sqlx.DB
}

func NewPostgreSQLRepository(db *sqlx.DB) models.AuthRepository {
	return &postgreSQLRepository{
		db: db,
	}
}

type DBUserWithProfile struct {
	UserID       uint      `db:"user_id"`
	ProfileID    uint      `db:"profile_id"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	IsAdmin      bool      `db:"is_admin"`
	Name         string    `db:"name"`
	DateOfBirth  time.Time `db:"dob"`
	Address      string    `db:"address"`
	Sex          string    `db:"sex"`
	PhoneNumber  string    `db:"phone_number"`
	ProfilePhoto string    `db:"profile_photo"`
}

func (ths *postgreSQLRepository) CreateUser(user models.User) (uint, uint, error) {
	dbUser := DBUserWithProfile{
		UserID:       user.ID,
		ProfileID:    user.Profile.ID,
		Email:        user.Email,
		Password:     user.Password,
		IsAdmin:      user.IsAdmin,
		Name:         user.Profile.Name,
		DateOfBirth:  user.Profile.DateOfBirth,
		Address:      user.Profile.Address,
		Sex:          user.Profile.Sex,
		PhoneNumber:  user.Profile.PhoneNumber,
		ProfilePhoto: user.Profile.ProfilePhoto,
	}

	rows, err := ths.db.NamedQuery(
		`WITH new_user AS (INSERT INTO "auth".user (email, password, is_admin) 
		VALUES (:email, :password, :is_admin) RETURNING id)
		INSERT INTO "auth".user_profile (name, dob, address, sex, phone_number, profile_photo, user_id) 
		VALUES (:name, :dob, :address, :sex, :phone_number, :profile_photo, (SELECT id FROM new_user)) 
		RETURNING user_id, id AS profile_id`,
		dbUser,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return 0, 0, models.NewClientError(
				"0100011",
				"user with email "+user.Email+" already exists",
				400,
			)
		}
		return 0, 0, models.NewServerError("0100011", 500, err)
	}

	var userID, profileID uint
	for rows.Next() {
		err = rows.Scan(&userID, &profileID)
		if err != nil {
			return 0, 0, models.NewServerError("0100012", 500, err)
		}
	}

	return userID, profileID, nil
}

func (ths *postgreSQLRepository) GetUserWithProfileByEmail(email string) (*models.User, error) {
	var dbUser DBUserWithProfile
	err := ths.db.Get(&dbUser,
		`SELECT u.id AS user_id, password, p.id AS profile_id
		FROM "auth".user AS u, "auth".user_profile AS p 
		WHERE u.id = p.user_id AND email=$1`,
		email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.NewClientError(
				"0100013",
				"no user found for email "+email,
				404,
			)
		}
		return nil, models.NewServerError("0100013", 500, err)
	}

	return &models.User{
		ID:       dbUser.UserID,
		Password: dbUser.Password,
		Profile: models.UserProfile{
			ID: dbUser.ProfileID,
		},
	}, nil
}

func (ths *postgreSQLRepository) GetUserWithProfileByID(userID uint) (*models.User, error) {
	var dbUser DBUserWithProfile
	err := ths.db.Get(&dbUser,
		`SELECT u.id AS user_id, email, p.id AS profile_id, name, 
		dob, address, sex, phone_number, profile_photo
		FROM "auth".user AS u, "auth".user_profile AS p 
		WHERE u.id = p.user_id AND u.id=$1`,
		userID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.NewClientError(
				"0100013",
				fmt.Sprintf("no user found for ID %d", userID),
				404,
			)
		}
		return nil, models.NewServerError("0100013", 500, err)
	}

	return &models.User{
		ID:      dbUser.UserID,
		Email:   dbUser.Email,
		IsAdmin: dbUser.IsAdmin,
		Profile: models.UserProfile{
			ID:           dbUser.ProfileID,
			Name:         dbUser.Name,
			DateOfBirth:  dbUser.DateOfBirth,
			Address:      dbUser.Address,
			Sex:          dbUser.Sex,
			PhoneNumber:  dbUser.PhoneNumber,
			ProfilePhoto: dbUser.ProfilePhoto,
		},
	}, nil
}
