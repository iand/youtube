/*
  This is free and unencumbered software released into the public domain. For more
  information, see <http://unlicense.org/> or the accompanying UNLICENSE file.
*/

// Client for the Youtube API
package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
}

type TextValue struct {
	Value string `json:"$t"`
}

func (t TextValue) String() string {
	return string(t.Value)
}

type Feed struct {
	ID      TextValue `json:"id"`
	Updated TextValue `json:"updated"`
	Title   TextValue `json:"title"`
	Logo    TextValue `json:"logo"`
	Links   []Link    `json:"link"`
	Entries []Entry   `json:"entry"`
}

type Link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type Entry struct {
	ID            TextValue     `json:"id"`
	Published     TextValue     `json:"published"`
	Updated       TextValue     `json:"updated"`
	Title         TextValue     `json:"title"`
	Links         []Link        `json:"link"`
	Media         MediaGroup    `json:"media$group"`
	Rating        Rating        `json:"gd$rating"`
	YoutubeRating YoutubeRating `json:"yt$rating"`
	Statistics    Statistics    `json:"yt$statistics""`
}

type Author struct {
	Name   TextValue `json:"name"`
	URI    TextValue `json:"uri"`
	UserID TextValue `json:"yt$userId"`
}

type MediaGroup struct {
	Thumbnails []Thumbnail `json:"media$thumbnail"`
	Duration   Duration    `json:"yt$duration"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Time   string `json:"time"`
	Name   string `json:"yt$name"`
}

type Duration struct {
	Seconds int `json:"seconds,string"`
}

type Rating struct {
	Average   float64 `json:"average"`
	Max       int     `json:"max"`
	Min       int     `json:"min"`
	NumRaters int     `json:"numRaters"`
}

type Statistics struct {
	FavoriteCount int `json:"favoriteCount,string"`
	ViewCount     int `json:"viewCount,string"`
}

type YoutubeRating struct {
	NumLikes    int `json:"numLikes,string"`
	NumDislikes int `json:"numDislikes,string"`
}

func New() *Client {
	return &Client{}
}

func (client *Client) VideoSearch(srch string) (*Feed, error) {
	var data struct {
		Feed Feed `json:"feed"`
	}

	url := fmt.Sprintf("https://gdata.youtube.com/feeds/api/videos?v=2&alt=json&q=%s", url.QueryEscape(srch))
	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("ArtistInfoByName failed with http error: %s", err.Error())
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&data); err != nil {
		return nil, fmt.Errorf("ArtistInfoByName failed to parse JSON response: %s", err.Error())
	}

	if err != nil {
		return nil, err
	}
	return &data.Feed, nil
}
