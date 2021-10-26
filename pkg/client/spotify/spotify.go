package spotify

type Spotify struct {
	clientID string
}

// New creates a new Spotify client.
func New(clientID string) *Spotify {
	return &Spotify{
		clientID: clientID,
	}
}
