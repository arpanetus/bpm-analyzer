package lastfm

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var testApiKey string

const Sixty = 60

func TestMain(t *testing.M) {
	log.Println("Initializing Last.fm API")

	testApiKey = os.Getenv("LASTFM_API_KEY")

	exitVal := t.Run()

	log.Println("Last.fm API test finished")
	
	os.Exit(exitVal)
}

func lastfm(t *testing.T) *LastFm {

	var c = &http.Client{
		Timeout: time.Second * Sixty,
	}
	
	defLog, errLog := log.New(os.Stdout, "[INFO] ", log.LstdFlags), log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	api, err := New(BaseURL, testApiKey, c, errLog, defLog)

	if err != nil {
		t.Errorf("Error creating Last.fm API: %s", err)
		t.FailNow()
	}

	return api
}