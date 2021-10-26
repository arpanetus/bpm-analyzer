package lastfm

import (
	"context"
	"reflect"
	"testing"
)

func TestLastFm_GetUserInfo(t *testing.T) {

	api := lastfm(t)

	u := UserInfo{User: User{
		Country:    "Falkland Islands (Malvinas)",
		Age:        "0",
		Playcount:  "4207",
		Subscriber: "0",
		Realname:   "2",
		Playlists:  "0",
		Bootstrap:  "0",
		Image:      []Image{{Size: Small, Text: ""}, {Size: Medium, Text: ""}, {Size: Large, Text: ""}, {Size: ExtraLarge, Text: ""}},
		Registered: Registered{Unixtime: "1043366400", Text: 1043366400},
		URL:        "https://www.last.fm/user/test",
		Gender:     "n",
		Name:       "test",
		Type:       "user"},
	}

	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserInfo
		wantErr bool
	}{
		{"get-user-test", args{context.Background(), "test"}, &u, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.GetUserInfo(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastFm.GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastFm.GetUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastFm_GetRecentTracks(t *testing.T) {

	api := lastfm(t)

	tr := RecentTracks{
		RecentTracks: RecentTracks_{
			Track: []Track{
				{
					Artist:     MbIDed{Mbid: "b7ffd2af-418f-4be2-bdd1-22f8b48613da", Text: "Nine Inch Nails"},
					Streamable: "0",
					Image:      []Image{{Size: Small, Text: ""}, {Size: Medium, Text: ""}, {Size: Large, Text: ""}, {Size: ExtraLarge, Text: ""}},
					MbID:       "05b88c1c-a9f5-39e1-ae32-680acb718260",
					Album:      MbIDed{Mbid: "", Text: ""},
					Name:       "The Line Begins to Blur",
					URL:        "https://www.last.fm/music/Nine+Inch+Nails/_/The+Line+Begins+to+Blur",
					Date:       Date{UTS: "1114653332", Text: "28 Apr 2005, 01:55"},
				},
			},
			Attr: Attr{
				User:       "test",
				TotalPages: "4207",
				Page:       "1",
				PerPage:    "1",
				Total:      "4207",
			},
		},
	}

	type args struct {
		ctx         context.Context
		username    string
		page, limit uint64
	}
	tests := []struct {
		name    string
		args    args
		want    *RecentTracks
		wantErr bool
	}{
		{"get-recent-tracks-test", args{context.Background(), "test", 1, 1}, &tr, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.GetRecentTracks(tt.args.ctx, tt.args.username, tt.args.page, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("LastFm.GetRecentTracks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastFm.GetRecentTracks() = %v, want %v", got, tt.want)
			}
		})
	}
}
