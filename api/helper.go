package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()


func (app *Config) ValidateBody(c *gin.Context, obj any) error {
	if err := c.BindJSON(&obj); err != nil {
		return err
	}

	if err := validate.Struct(&obj); err != nil {
		return err
	}

	return nil
}

func (app *Config) writeJSON(c *gin.Context, status int, data any) {
	c.JSON(status, jsonResponse{Status: status, Message: "success", Data: data})
}


func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	s := http.StatusBadRequest
	if len(status) > 0 {
		s = status[0]
	}
	payload := jsonResponse{
		Status:  s,
		Message: err.Error(),
		Data:    nil,
	}

	c.JSON(s, payload)
}