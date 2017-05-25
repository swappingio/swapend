package config

import (
	"log"

	"github.com/pajlada/kkonfig"
)

type Specification struct {
	Database struct {
		Hostname string `json:"Hostname"`
		Username string `json:"Username"`
		Password string `json:"Password"`
		Dbname   string `json:"Dbname"`
	} `json:"Database"`
	Mailgun struct {
		Enabled      bool   `json:"Enabled"`
		APIKey       string `json:"APIKey"`
		Domain       string `json:"Domain"`
		PublicAPIKey string `json:"PublicAPIKey"`
		URL          string `json:"URL"`
	} `json:"Mailgun"`
	Sessions struct {
		CookieSecret string `json:"CookieSecret"`
		StorageName  string `json:"StorageName"`
	} `json:"Authentication"`
	Redis struct {
		Hostname string `json:"Hostname"`
		Port     int    `json:"Port"`
		Timeout  int    `json:"Timeout"`
	} `json:"Redis"`
	Transcoder struct {
		ConcurrentTranscodes int  `json:"ConcurrentTranscodes"`
		Threads              int  `json:"Threads"`
		Debug                bool `json:"Debug"`
	} `json:"Transcoder"`
}

var (
	s Specification
)

func init() {

	var test []string
	test = append(test, "/home/coral/.streamcred.json")
	err := kkonfig.Process("web", test, &s)
	if err != nil {
		log.Fatal(err)
	}

	//Standard Values
	s.Transcoder.ConcurrentTranscodes = 1
	s.Transcoder.Threads = 1
	s.Transcoder.Debug = true

	log.Println("Loaded Config.")
}

func GetConfig() Specification {
	return s
}
