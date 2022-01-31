package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/erizkiatama/rasalibrary-backend/models"
)

func GenerateTokenPair(userID uint) (*models.TokenPair, error) {
	access := jwt.New(jwt.SigningMethodHS256)

	atClaims := access.Claims.(jwt.MapClaims)
	atClaims["access"] = true
	atClaims["userID"] = userID
	atClaims["expires"] = time.Now().Add(24 * time.Hour).Unix()

	// TODO: use env variable
	at, err := access.SignedString([]byte("rasalibrary-secret-key"))
	if err != nil {
		return nil, models.NewServerError("0100003", 500, err)
	}

	refresh := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refresh.Claims.(jwt.MapClaims)
	rtClaims["refresh"] = true
	rtClaims["userID"] = userID
	rtClaims["expires"] = time.Now().Add(24 * time.Hour).Unix()

	rt, err := refresh.SignedString([]byte("rasalibrary-secret-key"))
	if err != nil {
		return nil, models.NewServerError("0100004", 500, err)
	}

	return &models.TokenPair{
		Access:  at,
		Refresh: rt,
	}, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, models.NewServerError(
				"0100005",
				500,
				fmt.Errorf("unexpected signing method: %v", token.Header["alg"]),
			)

		}
		return []byte("rasalibrary-secret-key"), nil
	})
}
