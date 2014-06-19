package lastfm

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)

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



func NewLastfm(LastfmUser, LastfmKey string) *Lastfm {
	var track *Lastfm

	url := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=" + LastfmUser + "&api_key=" + LastfmKey + "&format=json"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &track); err != nil {
		log.Fatal(err)
	}
	return track
}


func (fm *Lastfm) IsNowPlaying() bool {
	return fm.Recenttracks.Track[0].Attr.Nowplaying != ""
}

func (fm *Lastfm) GetCurrentArtistAndTrackName() (string, string) {
	return fm.Recenttracks.Track[0].Artist.Text, fm.Recenttracks.Track[0].Name
}

func (fm *Lastfm) GetLastPlayedDate() string {
	val, err := strconv.ParseInt(fm.Recenttracks.Track[0].Date.Uts, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	tracktime := time.Unix(val, 0).UTC()
	lastPlay := time.Since(tracktime).String()
	return lastPlay
}
