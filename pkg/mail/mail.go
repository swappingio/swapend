package mail

import (
	"fmt"
	"log"

	"github.com/swappingio/swapend/pkg/config"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

func SendActivationEmail(username string, email string, verificationCode string) {
	s := config.GetConfig()
	if s.Mailgun.Enabled {
		mg := mailgun.NewMailgun(s.Mailgun.Domain, s.Mailgun.APIKey, s.Mailgun.PublicAPIKey)
		message := mailgun.NewMessage(
			"noreply@swapping.io",
			"Please verify your account at SWAPPING.IO",
			"Hello "+username+". \n Please verify your account at swapping.io by going here http://localhost:4020/verify/"+verificationCode,
			email)
		resp, id, err := mg.Send(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s Resp: %s\n", id, resp)
	}
}
