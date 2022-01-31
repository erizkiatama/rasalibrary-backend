package repository

import (
	"database/sql"
	"errors"
	"strings"

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

func (ths *postgreSQLRepository) CreateUser(user *models.User) error {
	rows, err := ths.db.NamedQuery(
		`INSERT INTO "auth".user (email, password, is_admin) 
		VALUES (:email, :password, :is_admin) RETURNING *`,
		user,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return models.NewClientError(
				"0100011",
				"user with email "+user.Email+" already exists",
				400,
			)
		} else {
			return models.NewServerError("0100011", 500, err)
		}
	}
	for rows.Next() {
		err = rows.StructScan(user)
		if err != nil {
			return models.NewServerError("0100012", 500, err)
		}
	}
	return nil
}

func (ths *postgreSQLRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ths.db.Get(&user, `SELECT * FROM "auth".user WHERE email=$1`, email)
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

	return &user, nil
}
