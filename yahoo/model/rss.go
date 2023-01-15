package model

import (
	"encoding/xml"
	"time"
)

type (
	// YahooNewsImage represents Yahoo news RSS
	YahooNewsRSS struct {
		XMLName xml.Name            `xml:"rss"`
		Version string              `xml:"version,attr"`
		Channel YahooNewsRSSChannel `xml:"channel"`
	}

	// YahooNewsRSSChannel represents Yahoo news RSS Channel tag
	YahooNewsRSSChannel struct {
		Language       string             `xml:"language"`
		Copyright      string             `xml:"copyright"`
		PubDateStr     string             `xml:"pubDate"`
		Title          string             `xml:"title"`
		Link           string             `xml:"link"`
		Description    string             `xml:"description"`
		Image          YahooNewsRSSImage  `xml:"image"`
		YahooNewsItems []YahooNewsRSSItem `xml:"item"`
	}

	// YahooNewsRSSImage represents Yahoo news RSS image tag
	YahooNewsRSSImage struct {
		Title  string `xml:"title"`
		Link   string `xml:"link"`
		URL    string `xml:"url"`
		Width  int64  `xml:"width"`
		Height int64  `xml:"height"`
	}

	// YahooNewsImage represents Yahoo news RSS item tag
	YahooNewsRSSItem struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		PubDateStr  string `xml:"pubDate"`
		Image       string `xml:"image"`
		Comments    string `xml:"comments"`
		Description string `xml:"description"`
	}
)

func (y *YahooNewsRSS) Items() []YahooNewsRSSItem {
	return y.Channel.YahooNewsItems
}

func (y *YahooNewsRSSChannel) PubDate() (time.Time, error) {
	return time.Parse(time.RFC1123, y.PubDateStr)
}

func (y *YahooNewsRSSItem) PubDate() (time.Time, error) {
	return time.Parse(time.RFC1123, y.PubDateStr)
}
