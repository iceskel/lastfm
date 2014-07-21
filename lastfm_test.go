package lastfm

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"testing"
)

type configuration struct {
	Channel               string
	Botname               string
	Aouth                 string
	LastfmKey             string
	LastfmSecret          string
	LastfmUser            string
	RepeatMsg             string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
	TwitterAccessToken    string
	TwitterAccessSecret   string
}

var (
	config configuration
)

func Test_NewLastFm(t *testing.T) {
	_, err := NewLastfm(config.LastfmUser, config.LastfmKey)
	if err != nil {
		t.Error("Failed")
	}

}

func Test_GetLastPlayedDate(t *testing.T) {
	fm, err := NewLastfm(config.LastfmUser, config.LastfmKey)
	if err != nil {
		t.Error("Failed")
	}
	if _, err := fm.GetLastPlayedDate(); err != nil {
		t.Error("Failed")
	}
}

func init() {
	configFile := flag.String("c", "conf.json", "config file")
	flag.Parse()

	file, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}
}
