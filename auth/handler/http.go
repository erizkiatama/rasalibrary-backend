package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/erizkiatama/rasalibrary-backend/auth/helper"
	"github.com/erizkiatama/rasalibrary-backend/models"
	"github.com/gin-gonic/gin"
)

type handler struct {
	svc models.AuthService
}

func NewAuthHandler(router *gin.RouterGroup, svc models.AuthService) {
	handler := &handler{
		svc: svc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
		auth.POST("/login/refresh", handler.RefreshToken)
	}
}

func (ths *handler) Login(c *gin.Context) {
	var req models.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		models.ConstructResponse(c, nil, models.NewClientError(
			"0100008",
			"invalid parameters: "+err.Error(),
			400,
		))
		return
	}

	token, err := ths.svc.Login(req)
	if err != nil {
		models.ConstructResponse(c, nil, err)
		return
	}

	models.ConstructResponse(c, token, nil)
}

func (ths *handler) Register(c *gin.Context) {
	var req models.RegisterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		models.ConstructResponse(c, nil, models.NewClientError(
			"0100009",
			"invalid parameters: "+err.Error(),
			400,
		))
		return
	}

	token, err := ths.svc.Register(req)
	if err != nil {
		models.ConstructResponse(c, nil, err)
		return
	}

	models.ConstructResponse(c, token, nil)
}

func (ths *handler) RefreshToken(c *gin.Context) {
	var tokenPair models.TokenPair

	c.ShouldBindJSON(&tokenPair)
	refreshToken, err := helper.ValidateToken(tokenPair.Refresh)
	if err != nil {
		models.ConstructResponse(c, nil, err)
		return
	}

	if rtClaims, ok := refreshToken.Claims.(jwt.MapClaims); ok && refreshToken.Valid {
		if rtClaims["refresh"] == true {
			userID := rtClaims["userID"].(float64)
			profileID := rtClaims["profileID"].(float64)
			newTokenPair, err := helper.GenerateTokenPair(uint(userID), uint(profileID))
			if err != nil {
				models.ConstructResponse(c, nil, err)
				return
			}
			newTokenPair.Refresh = ""
			models.ConstructResponse(c, newTokenPair, nil)
			return
		}
	}

	models.ConstructResponse(c, nil, models.NewClientError(
		"0100010",
		"invalid refresh token",
		401,
	))
}
