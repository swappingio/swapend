package config

import (
	"fmt"

	"github.com/pajlada/kkonfig"
)

type Specification struct {
	Database struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Hostname string `json:"Hostname"`
		Dbname   string `json:"Dbname"`
	} `json:"Database"`
	Mailgun struct {
		APIKey       string `json:"APIKey"`
		Domain       string `json:"Domain"`
		PublicAPIKey string `json:"PublicAPIKey"`
		URL          string `json:"URL"`
	} `json:"Mailgun"`
	Authentication struct {
		CookieSecret string `json:"CookieSecret"`
		StorageName  string `json:"StorageName"`
	} `json:"Authentication"`
}

var (
	s Specification
)

func init() {

	var test []string
	test = append(test, "/home/coral/.streamcred.json")
	err := kkonfig.Process("web", test, &s)
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfig() Specification {
	return s
}
