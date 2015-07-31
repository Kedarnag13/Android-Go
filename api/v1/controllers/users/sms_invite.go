package users

import (
	"github.com/sfreiberg/gotwilio"
	"net/http"
)

type messageController struct{}

var Message messageController

func (m messageController) Send(rw http.ResponseWriter, req *http.Request) {
	accountSid := "AC62e6cfa0a301115e75e91cd6fb176e5f"
	authToken := "d25a1dd728cec4edc7d1b1e34d9b04b6"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+14805683241"
	to := "+917022665448"
	message := "Hello, Welcome to Android-Go. Please click on the following link to download the application. " + "http://google.com"
	twilio.SendSMS(from, to, message, "", "")
}
