package api

import{
	"errors"

	"github.com/twilio/twilio-go"
	"github.com/twilio/twilio-go/rest/verify/v2"
}
var client *twilio.RestClient=twilio.NewRestClientWithParams(twilio.ClientParams{
	Username:envACCOUNTSID(),
	Password:envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string,error) {
	params:=&verify.CreateVerificationParams{}
	params.To(phoneNumber)
	params.SetChannel("sms")
	res,err:=client.VerifyV2.CreateVerification(envSERVICESID(),params)
	if err!=nil{
		return "",err
	}
	return *res.Sid,nil
}

func (app *Config) twilioVerifyOTP(phoneNumber string,code string) error{
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)
	res,err:=client.VerifyV2.CreateVerificationCheck(envSERVICESID(),params)
	if err!=nil{
		return err
	}
	return nil
}