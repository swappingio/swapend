package mail

import (
	"fmt"
	"log"

	"github.com/coral/swapend/pkg/config"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

func SendVerificationEmail(username string, email string, verificationCode string) {
	s := config.GetConfig()

	fmt.Println(username, email, verificationCode)
	mg := mailgun.NewMailgun(s.Mailgun.Domain, s.Mailgun.APIKey, s.Mailgun.PublicAPIKey)
	message := mailgun.NewMessage(
		"noreply@swapping.io",
		"Please verify your account at SWAPPING.IO",
		"Hello "+username+". Please verify your account, here is a string"+verificationCode,
		email)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
