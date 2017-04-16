package config

import (
	"fmt"

	"github.com/pajlada/kkonfig"
)

type Specification struct {
	Database struct {
		Username string `default:"coral"`
		Password string `default:"lolwut"`
		Hostname string `default:"localhost"`
		Dbname   string `default:"lolwut"`
	} `json:"Database"`
	Mailgun struct {
		APIKey       string `json:"APIKey"`
		Domain       string `json:"Domain"`
		PublicAPIKey string `json:"PublicAPIKey"`
		URL          string `json:"URL"`
	} `json:"Mailgun"`
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
