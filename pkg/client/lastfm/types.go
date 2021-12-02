package lastfm

type Size string

const (
	Small      Size = "small"
	Medium     Size = "medium"
	Large      Size = "large"
	ExtraLarge Size = "extralarge"
)

type Image struct {
	Size Size   `json:"size"`
	Text string `json:"#text"`
}

type Registered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

type User struct {
	Country    string     `json:"country"`
	Age        string     `json:"age"`
	Playcount  string     `json:"playcount"`
	Subscriber string     `json:"subscriber"`
	Realname   string     `json:"realname"`
	Playlists  string     `json:"playlists"`
	Bootstrap  string     `json:"bootstrap"`
	Image      []Image    `json:"image"`
	Registered Registered `json:"registered"`
	URL        string     `json:"url"`
	Gender     string     `json:"gender"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
}

type UserInfo struct {
	User User `json:"user"`
}

type MbIDed struct {
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}

type Date struct {
	UTS  string `json:"uts"`
	Text string `json:"#text"`
}

type Track struct {
	Artist     MbIDed  `json:"artist"`
	Streamable string  `json:"streamable"`
	Image      []Image `json:"image"`
	MbID       string  `json:"mbid"`
	Album      MbIDed  `json:"album"`
	Name       string  `json:"name"`
	URL        string  `json:"url"`
	Date       Date    `json:"date"`
}

type Attr struct {
	User       string `json:"user"`
	TotalPages string `json:"totalPages"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	Total      string `json:"total"`
}

type RecentTracks_ struct {
	Tracks []*Track `json:"track"`
	Attr   Attr    `json:"@attr"`
}

type RecentTracks struct {
	RecentTracks RecentTracks_ `json:"recenttracks"`
}
