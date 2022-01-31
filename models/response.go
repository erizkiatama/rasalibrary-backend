package models

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  error       `json:"error,omitempty"`
}

func ConstructResponse(c *gin.Context, result interface{}, err error) {
	var response Response
	var customErr CustomError

	if err != nil {
		response.Error = err
		if ok := errors.As(err, &customErr); !ok {
			return
		}
		customErr.LogError()
		c.AbortWithStatusJSON(customErr.GetStatus(), response)
		return
	}

	response.Result = result
	c.JSON(http.StatusOK, response)
}
