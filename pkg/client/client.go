package client

import (
	"context"
	"github.com/arpanetus/bpm-analyzer/pkg/client/lastfm"
)

type LastFmer interface {
	GetUserInfo(ctx context.Context, username string) (*lastfm.UserInfo, error)
	GetRecentTracks(
		ctx context.Context,
		username string,
		page, limit uint64,
	) (*lastfm.RecentTracks, error)
}

type Spotifier interface {
}
