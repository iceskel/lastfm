package lastfm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Lastfm represents the Recenttracks json
// api.
type Lastfm struct {
	Recenttracks struct {
		Track []struct {
			Artist struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"artist"`
			Streamable string `json:"streamable"`
			Album      struct {
				Text string `json:"#text"`
				Mbid string `json:"mbid"`
			} `json:"album"`
			Image []struct {
				Text string `json:"#text"`
				Size string `json:"size"`
			} `json:"image"`
			Name string `json:"name"`
			Mbid string `json:"mbid"`
			Url  string `json:"url"`
			Attr struct {
				Nowplaying string `json:"nowplaying"`
			} `json:"@attr"`
			Date struct {
				Text string `json:"#text"`
				Uts  string `json:"uts"`
			} `json:"date"`
		} `json:"track"`
		Attr struct {
			User       string `json:"user"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"recenttracks"`
}

// NewLastfm returns a new Lastfm instance with the given
// LastfmUser and LastfmKey.
func NewLastfm(LastfmUser, LastfmKey string) (*Lastfm, error) {
	var track *Lastfm

	url := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + LastfmUser + "&api_key=" + LastfmKey + "&format=json"
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &track); err != nil {
		return nil, err
	}
	return track, nil
}

// IsNowPlaying returns the current track
// that is playing.
func (fm *Lastfm) IsNowPlaying() bool {
	return fm.Recenttracks.Track[0].Attr.Nowplaying != ""
}

// GetCurrentArtistAndTrackName returns the current artist name
// and track name of the track that is currently playing.
func (fm *Lastfm) GetCurrentArtistAndTrackName() (string, string) {
	return fm.Recenttracks.Track[0].Artist.Text, fm.Recenttracks.Track[0].Name
}

// GetLastPlayedDate returns the date of the last song that
// was played.
func (fm *Lastfm) GetLastPlayedDate() (string, error) {
	val, err := strconv.ParseInt(fm.Recenttracks.Track[0].Date.Uts, 10, 64)
	if err != nil {
		return "", err
	}
	tracktime := time.Unix(val, 0).UTC()
	return time.Since(tracktime).String(), nil
}
