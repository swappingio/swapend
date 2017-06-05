package config

import (
	"github.com/pajlada/kkonfig"
	"log"
	"os"
	"runtime"
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

	// will clean this up later, but for now it's okay 
	if runtime.GOOS == "windows"{
		pathSeperator = "\\"

	} else if runtime.GOOS == "darwin"{
		pathSeperator = "/"

	} else if runtime.GOOS == "linux" {
		pathSeperator = "/"

	} else {
		log.Printf("THIS IS UNACEPTABLE !!")
		os.Exit(0) //something happend that wasn't supposed to happend...
	}

	var configFile string
	configFile = "/.streamcred.json"

	usr, hpath := user.Current()
	if hpath != nil {
		log.Fatal(hpath)
	}

	var test []string
	test = append(test, usr.HomeDir+configFile)
	err := kkonfig.Process("web", test, &s)
	if err != nil {
		log.Fatal(err)
	}

	PATH := usr.HomeDir+pathSeperator+configFile

	_, fileErr := os.Stat(PATH)
	if os.IsNotExist(fileErr){
		log.Printf("The config '%s' doesn't exist.\n", configFile)
		os.Exit(127) // bash exit code for problems with path
	} else {
		log.Printf("\nFound Config. \n%s", PATH)
	}


	//Standard Values
	s.Transcoder.ConcurrentTranscodes = 1
	s.Transcoder.Threads = 1
	s.Transcoder.Debug = true

	//log.Println("Loaded Config.")
}

func GetConfig() Specification {
	return s
}
