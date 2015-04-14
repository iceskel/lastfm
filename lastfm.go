package lastfm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// lastfmData represents the Recenttracks json
// api.
type lastfmData struct {
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

// LastfmApi struct for the api.
type LastfmApi struct {
	LastfmUser string
	LastfmKey  string
}

// New returns a new Lastfm instance with the given
// LastfmUser and LastfmKey.
func New(lastfmUser, lastfmKey string) *LastfmApi {
	return &LastfmApi{
		LastfmUser: lastfmUser,
		LastfmKey:  lastfmKey,
	}
}

// IsNowPlaying returns the current track
// that is playing.
func (fm *LastfmApi) IsNowPlaying() (bool, error) {
	var track *lastfmData
	url := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + fm.LastfmUser +
		"&api_key=" + fm.LastfmKey + "&format=json"
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(body, &track); err != nil {
		return false, err
	}

	return track.Recenttracks.Track[0].Attr.Nowplaying != "", nil
}

// GetCurrentArtistAndTrackName returns the current artist name
// and track name of the track that is currently playing.
func (fm *LastfmApi) GetCurrentArtistAndTrackName() (string, string, error) {
	var track *lastfmData
	url := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + fm.LastfmUser +
		"&api_key=" + fm.LastfmKey + "&format=json"
	res, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	if err := json.Unmarshal(body, &track); err != nil {
		return "", "", err
	}

	return track.Recenttracks.Track[0].Artist.Text, track.Recenttracks.Track[0].Name, nil
}

// GetLastPlayedDate returns the date of the last song that
// was played.
func (fm *LastfmApi) GetLastPlayedDate() (string, error) {
	var track *lastfmData
	url := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + fm.LastfmUser +
		"&api_key=" + fm.LastfmKey + "&format=json"
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &track); err != nil {
		return "", err
	}

	val, err := strconv.ParseInt(track.Recenttracks.Track[0].Date.Uts, 10, 64)
	if err != nil {
		return "", err
	}
	tracktime := time.Unix(val, 0).UTC()

	return time.Since(tracktime).String(), nil
}
