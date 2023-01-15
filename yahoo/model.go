package yahoo

import (
	"encoding/xml"
	"time"
)

type (
	YahooNewsRSS struct {
		XMLName xml.Name         `xml:"rss"`
		Version string           `xml:"version,attr"`
		Channel YahooNewsChannel `xml:"channel"`
	}

	YahooNewsChannel struct {
		Language       string          `xml:"language"`
		Copyright      string          `xml:"copyright"`
		PubDateStr     string          `xml:"pubDate"`
		Title          string          `xml:"title"`
		Link           string          `xml:"link"`
		Description    string          `xml:"description"`
		Image          YahooNewsImage  `xml:"image"`
		YahooNewsItems []YahooNewsItem `xml:"item"`
	}

	YahooNewsImage struct {
		Title  string `xml:"title"`
		Link   string `xml:"link"`
		URL    string `xml:"url"`
		Width  int64  `xml:"width"`
		Height int64  `xml:"height"`
	}

	YahooNewsItem struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		PubDateStr  string `xml:"pubDate"`
		Image       string `xml:"image"`
		Comments    string `xml:"comments"`
		Description string `xml:"description"`
	}
)

func (y *YahooNewsChannel) PubDate() (time.Time, error) {
	return time.Parse(time.RFC1123, y.PubDateStr)
}

func (y *YahooNewsItem) PubDate() (time.Time, error) {
	return time.Parse(time.RFC1123, y.PubDateStr)
}
