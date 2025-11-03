package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karthikbhandary2/otp/data"
)

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*10)
		var payload data.OTPData
		defer cancel()

		app.ValidateBody(c, &payload)

		newData := data.OTPData{
			Phone: payload.Phone,
		}

		_, err := app.twilioSendOTP(newData.Phone)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var payload data.VerifyData
		defer cancel()

		app.ValidateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.Phone, newData.Code)
		fmt.Println("err: ", err)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}