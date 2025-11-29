package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imshubham024/emailVerify/data"
)

const appTimeout=time.Second*10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payLoad data.OTPData
		app.validateBody(c, &payLoad)
		defer cancel()
		newData := data.OTPData{
			PhoneNumber: payLoad.PhoneNumber,
		}
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errJSON(c, err)
			return
		}
		app.writeJSON
	}
}
