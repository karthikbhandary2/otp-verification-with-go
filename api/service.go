package api

import (
	"errors"
	"github.com/twilio/twilio-go"
	twilioApi"github.com/twilio/twilio-go/rest/verify/v2"
)

var TwilioClient *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envACCOUNTSID(),
	Password: envAUTHTOKEN(),
})

func (app *Config) twilioSendOTP(phone string) (string,error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phone)
	params.SetChannel("sms")

	resp, err := TwilioClient.VerifyV2.CreateVerification(envSERVICESID(),params)
	if err != nil {
		return "",err
	}

	return *resp.Sid,nil

}

func (app *Config) twilioVerifyOTP(phone string,code string) (error) {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phone)
	params.SetCode(code)

	resp, err := TwilioClient.VerifyV2.CreateVerificationCheck(envSERVICESID(),params)
	if err != nil {
		return err
	}

	if *resp.Status != "approved" {
		return errors.New("not a valid code")
	}

	return nil
}