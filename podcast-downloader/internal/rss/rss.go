package rss

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Enclosure   struct {
		URL string `xml:"url,attr"`
	} `xml:"enclosure"`
}

type RSSFeed struct {
	Channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Items       []Item `xml:"item"`
	} `xml:"channel"`
}

func FetchRSSFeed(url string) (*RSSFeed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch RSS feed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	return &feed, nil
}
