package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imshubham024/go-emailVerify/data"
)

const appTimeout = time.Second * 10

func (app *Config) sendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var payLoad data.OTPData

		if err := app.validateBody(c, &payLoad); err != nil {
			app.errJSON(c, err)
			return 
		}

		newData := data.OTPData{
			PhoneNumber: payLoad.PhoneNumber,
		}
		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "SMS sent successfully")
	}
}

func (app *Config) verifySMS() gin.HandlerFunc{
	return func(c *gin.Context) {
		_,cancel:=context.WithTimeout(context.Background(),appTimeout)
		defer cancel()
		var payload data.VerifyOTP
		
		if err:=app.validateBody(c,&payload);err!=nil{
			app.errJSON(c,err)
			return
		}

		newData:=data.VerifyOTP{
			User:payload.User,
			Code:payload.Code,
		}
		err:=app.twilioVerifyOTP(newData.User.PhoneNumber,newData.Code)
		if err!=nil{
			app.errJSON(c,err)
			return
		}
		app.writeJSON(c,http.StatusAccepted,"OTP verified successfully")
	}
}
