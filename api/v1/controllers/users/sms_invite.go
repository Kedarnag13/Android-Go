package users

import (
	"fmt"
	"github.com/njern/gonexmo"
	"log"
	"net/http"
)

type messageController struct{}

var Message messageController

func (m messageController) Send(rw http.ResponseWriter, req *http.Request) {
	nexmo_client, _ := nexmo.NewClientFromAPI("4fc94b4e", "eb628653")

	// Test if it works by retrieving your account balance
	balance, err := nexmo_client.Account.GetBalance()
	if err != nil || balance == 0.0 {
		log.Fatal(err)
	}
	// Send an SMS
	// See https://docs.nexmo.com/index.php/sms-api/send-message for details.
	smsMsg := &nexmo.SMSMessage{
		From:     "+919916854300",
		To:       "+917022665448",
		Body:     []byte("Welcome to Android-Go"),
		Type:     nexmo.WAPPush,
		Title:    "Hello World",
		URL:      "http://www.google.com",
		Validity: 1000,
	}

	messageResponse, err := nexmo_client.SMS.Send(smsMsg)
	if err != nil || messageResponse == nil {
		log.Fatal(err)
	}
	fmt.Println(messageResponse)
}
