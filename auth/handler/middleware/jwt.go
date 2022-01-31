package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/erizkiatama/rasalibrary-backend/auth/helper"
	"github.com/erizkiatama/rasalibrary-backend/models"
	"github.com/gin-gonic/gin"
)

func AuthorizeTokenJWT(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		models.ConstructResponse(c, nil, models.NewClientError(
			"0100006",
			"authorization header not given",
			401,
		))
		return
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := helper.ValidateToken(tokenString)
	if err != nil {
		models.ConstructResponse(c, nil, err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(float64)
		profileID := claims["profileID"].(float64)
		c.Set("userID", uint(userID))
		c.Set("profileID", uint(profileID))
		c.Next()
	} else {
		models.ConstructResponse(c, nil, models.NewClientError(
			"0100007",
			"invalid token",
			401,
		))
	}
}
