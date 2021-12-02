package main

import (
	"context"
	"github.com/arpanetus/bpm-analyzer/pkg/client/lastfm"
	"github.com/arpanetus/bpm-analyzer/pkg/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()

	var c = &http.Client{
		Timeout: time.Second * 60,
	}

	defLog, errLog := log.New(os.Stdout, "[INFO] ", log.LstdFlags), log.New(os.Stderr, "[ERROR] ", log.LstdFlags)

	api, err := lastfm.New(
		lastfm.BaseURL,
		"c364eb4d7be80d53b2532abc191b8397",
		c, errLog, defLog,
	)

	if err != nil {
		defLog.Fatalf("cannot create lastfm client: %+v", err)
	}

	svc := service.NewLastFmService(api, errLog, defLog)

	if err = svc.DownloadIntoFile(ctx, "arpanetus", "/tmp/lastfm"); err != nil {
		defLog.Fatalf("cannot fetch data from lastfm: %+v", err)
	}
}
