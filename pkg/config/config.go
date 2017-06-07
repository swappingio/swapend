package config

import (
	"log"
	"os"
	"os/user"
	"runtime"

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

	var pathSeperator string

	if runtime.GOOS == "windows" {
		pathSeperator = "\\"

	} else {
		pathSeperator = "/"
	}

	var configFile string
	configFile = ".streamcred.json"

	usr, hpath := user.Current()
	if hpath != nil {
		log.Fatal(hpath)
	}

	PATH := usr.HomeDir + pathSeperator + configFile

	var test []string
	test = append(test, PATH)
	err := kkonfig.Process("web", test, &s)
	if err != nil {
		log.Fatal(err)
	}

	_, fileErr := os.Stat(PATH)
	if os.IsNotExist(fileErr) {
		log.Printf("The config '%s' doesn't exist.\n", configFile)
		os.Exit(127)
	} else {
		log.Printf("\nFound Config. \n%s", PATH)
	}

	//Standard Values
	s.Transcoder.ConcurrentTranscodes = 1
	s.Transcoder.Threads = 1
	s.Transcoder.Debug = true
}

func GetConfig() Specification {
	return s
}
