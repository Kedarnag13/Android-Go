package users

import (
	"fmt"
	"github.com/njern/gonexmo"
	"log"
	"net/http"
	"strconv"
	"time"
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
		From:            "+919916854300",
		To:              "+917022665448",
		Type:            nexmo.Text,
		Text:            "Hello, Welcome to Android-Go. Please click on the following link to download the application. " + "http://google.com",
		ClientReference: "gonexmo-test " + strconv.FormatInt(time.Now().Unix(), 10),
		Class:           nexmo.Flash,
	}

	messageResponse, err := nexmo_client.SMS.Send(smsMsg)
	if err != nil || messageResponse == nil {
		log.Fatal(err)
	}
	fmt.Println(messageResponse)
}
