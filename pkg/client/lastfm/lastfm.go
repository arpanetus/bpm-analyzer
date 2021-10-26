package lastfm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type LastFm struct {
	url            url.URL
	apiKey         string
	client         http.Client
	errLog, defLog *log.Logger
}

var (
	ErrUrlParse = errors.New("cannot parse base url")
)

const (
	BaseURL = "http://ws.audioscrobbler.com/2.0/"
	intBase = 10
)

// New creates a new Lastfm client.
func New(
	baseURL string,
	apiKey string,
	client *http.Client,
	errLog, defLog *log.Logger,
) (*LastFm, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		errLog.Printf("cannot parse base url: %s", err)

		return nil, ErrUrlParse
	}

	q := u.Query()
	q.Add("format", "json")
	u.RawQuery = q.Encode()

	return &LastFm{
		url:    *u,
		apiKey: apiKey,
		client: *client,
		errLog: errLog,
		defLog: defLog,
	}, nil
}

var (
	getUserInfoPfx     = "[GetUserInfo]: "
	getRecentTracksPfx = "[GetRecentTracks]: "
)

// GetUserInfo returns the user info.
func (l *LastFm) GetUserInfo(ctx context.Context, username string) (*UserInfo, error) {
	l.defLog.Print(getUserInfoPfx + username)

	url := l.url
	q := url.Query()
	q.Add("method", "user.getInfo")
	q.Add("user", username)
	q.Add("api_key", l.apiKey)
	url.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		l.errLog.Printf(getUserInfoPfx+"cannot create GetUserInfo request: %s", err)

		return nil, fmt.Errorf(getUserInfoPfx+"cannot create request: %w", err)
	}

	resp, err := l.client.Do(req)
	if err != nil {
		l.errLog.Printf(getUserInfoPfx+"cannot get GetUserInfo response: %s", err)

		return nil, fmt.Errorf(getUserInfoPfx+"cannot get response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		l.errLog.Printf(getUserInfoPfx+"")

		return fmt.Errorf("unexpected status code: %", resp.StatusCode)
	}

	u := new(UserInfo)

	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		l.errLog.Printf(getUserInfoPfx+"cannot decode GetUserInfo response: %s", err)

		return nil, fmt.Errorf(getUserInfoPfx+"cannot decode response: %w", err)
	}

	return u, nil
}

func (l *LastFm) GetRecentTracks(
	ctx context.Context,
	username string,
	page, limit uint64,
) (*RecentTracks, error) {
	pageStr, limitStr := strconv.FormatUint(page, intBase), strconv.FormatUint(limit, intBase)
	l.defLog.Print(getRecentTracksPfx + username + ", " + pageStr + ", " + limitStr)

	url := l.url
	q := url.Query()
	q.Add("method", "user.getRecentTracks")
	q.Add("user", username)
	q.Add("api_key", l.apiKey)
	q.Add("page", pageStr)
	q.Add("limit", limitStr)

	url.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		l.errLog.Printf(getRecentTracksPfx+"cannot create GetUserInfo request: %s", err)

		return nil, fmt.Errorf(getRecentTracksPfx+"cannot create request: %w", err)
	}

	resp, err := l.client.Do(req)
	if err != nil {
		l.errLog.Printf(getRecentTracksPfx+"cannot get GetUserInfo response: %s", err)

		return nil, fmt.Errorf(getRecentTracksPfx+"cannot get response: %w", err)
	}

	u := new(RecentTracks)

	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		l.errLog.Printf(getRecentTracksPfx+"cannot decode GetUserInfo response: %s", err)

		return nil, fmt.Errorf(getRecentTracksPfx+"cannot decode response: %w", err)
	}

	return u, nil
}
