package lastfm

type LastFm struct {
	ApiKey string
}

// New creates a new Lastfm client.
func New(apiKey string) *LastFm {
	return &LastFm{
		ApiKey: apiKey,
	}
}

