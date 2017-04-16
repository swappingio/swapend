package mail

import (
	"fmt"
	"log"

	"github.com/coral/swapend/pkg/config"
	mailgun "gopkg.in/mailgun/mailgun-go.v1"
)

func TestMail() {
	s := config.GetConfig()
	mg := mailgun.NewMailgun(s.Mailgun.Domain, s.Mailgun.APIKey, s.Mailgun.PublicAPIKey)
	message := mailgun.NewMessage(
		"hello@swapping.io",
		"Here is song from ROUND 1",
		"ROUND 1",
		"someone@swapping.io")
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
