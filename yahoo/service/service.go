//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"context"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/tMinamii/gonews/yahoo/conf"
	"github.com/tMinamii/gonews/yahoo/fetcher"
	"github.com/tMinamii/gonews/yahoo/model"
)

type YahooNews interface {
	FetchAndStore(ctx context.Context) error
}

type YahooNewsService struct {
	RSSFetcher          fetcher.RSSFetcher
	HTMLDocumentFetcher fetcher.HTMLDocumentFetcher
}

func NewYahooNewsService() YahooNews {
	return &YahooNewsService{
		RSSFetcher:          fetcher.NewYahooNewsRSSFetcher(),
		HTMLDocumentFetcher: fetcher.NewWebsiteFetcher(),
	}
}

func (y *YahooNewsService) FetchAndStore(ctx context.Context) error {
	for _, v := range conf.YahooNewsAllGenre {
		rss, err := y.fetchRSS(ctx, v)
		if err != nil {
			continue
		}
		_ = y.fetchNewsHTML(ctx, rss)
	}
	return nil
}

func (y *YahooNewsService) fetchRSS(ctx context.Context, genre conf.YahooNewsGenre) (*model.YahooNewsRSS, error) {
	url := conf.YahooNewsMap[genre]
	rss, err := y.RSSFetcher.Fetch(ctx, url)
	if err != nil {
		return nil, err
	}
	return rss, nil
}

func (y *YahooNewsService) fetchNewsHTML(ctx context.Context, rss *model.YahooNewsRSS) []*goquery.Document {
	items := rss.Items()
	docs := make([]*goquery.Document, 0, len(items))
	for _, v := range items {
		doc, err := y.HTMLDocumentFetcher.Fetch(ctx, v.Link)
		if err != nil {
			log.Printf("failed to fetch website: title = %s url = %s", v.Title, v.Link)
			continue
		}
		docs = append(docs, doc)
	}
	return docs
}

func (y *YahooNewsService) scrapeContent(ctx context.Context, rss *goquery.Document) (*string, error) {
	return nil, nil
}
