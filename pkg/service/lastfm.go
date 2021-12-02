package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/arpanetus/bpm-analyzer/pkg/client"
	"github.com/arpanetus/bpm-analyzer/pkg/client/lastfm"
	"log"
	"os"
)

type LastFmService struct {
	client         client.LastFmer
	errLog, defLog *log.Logger
}

const limit uint64 = 200

func NewLastFmService(client client.LastFmer, errLog, defLog *log.Logger) *LastFmService {
	return &LastFmService{
		client: client,
		errLog: errLog,
		defLog: defLog,
	}
}

func (s *LastFmService) DownloadIntoFile(ctx context.Context, username, path string) (err error) {
	_, err = s.client.GetUserInfo(ctx, username)
	if err != nil {
		return fmt.Errorf("cannot get user info for %s : %w", username, err)
	}

	log.Printf("gonna fetch data for {%s}.", username)

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create file to fetch into: %w", err)
	}

	trackChan, errChan := make(chan *lastfm.Track, limit), make(chan error, 1)

	defer func() {
		if err != nil {
			closeErr := file.Close()
			s.errLog.Printf("cannot close the file, yet another occurred error: %w", closeErr)
		} else if err = file.Close(); err != nil {
			err = fmt.Errorf("cannot close the file: %w", err)
		}
		s.defLog.Print("finished fetching into file")
	}()

	defer close(trackChan)
	defer close(errChan)

	go func() {
		errChan <- s.fetcher(ctx, username, trackChan)
	}()

	for {
		select {
		case track := <-trackChan:
			data, err := json.Marshal(track)
			if err != nil {
				return fmt.Errorf("cannot serialize fetched track: %w", err)
			}

			data = append(data, []byte("\n")...)

			if _, err = file.Write(data); err != nil {
				return fmt.Errorf("cannot write a line into file: %w", err)
			}
		case err = <-errChan:
			if err != nil {
				return fmt.Errorf("cannot fetch data: %w", err)
			}

			return
		}
	}

	return nil
}

func (s *LastFmService) fetcher(ctx context.Context, username string, into chan<- *lastfm.Track) error {
	page, willFetch := uint64(1), true

	for willFetch {
		tracks, err := s.client.GetRecentTracks(ctx, username, page, limit)
		if err != nil {
			return fmt.Errorf("cannot fetch tracks: %w", err)
		}

		willFetch = len(tracks.RecentTracks.Tracks) != 0

		for _, t := range tracks.RecentTracks.Tracks {
			into <- t
		}

		page++
	}

	return nil
}

func (s *LastFmService) UniqueTracks(ctx context.Context, tracks []*lastfm.Track) map[string]*lastfm.Track {
	uniq := make(map[string]*lastfm.Track, len(tracks))
	for _, t := range tracks {
		h := t.Name + t.Album.Text + t.Artist.Text
		if _, ok := uniq[h]; !ok {
			uniq[h] = t
		}
	}

	return uniq
}
