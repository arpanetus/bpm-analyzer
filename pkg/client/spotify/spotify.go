package spotify

type Spotify struct {
	token string
}

// New creates a new Spotify client.
func New(token string) *Spotify {
	return &Spotify{
		token: token,
	}
}
