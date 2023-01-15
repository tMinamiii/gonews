//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package service

import (
	"context"

	"github.com/tMinamii/gonews/yahoo/conf"
	"github.com/tMinamii/gonews/yahoo/fetcher"
	"github.com/tMinamii/gonews/yahoo/model"
)

type YahooNews interface {
	FetchAndStore(ctx context.Context) error
}

type YahooNewsService struct {
	Fetcher fetcher.RSSFetcher
}

func NewYahooNewsService() YahooNews {
	return &YahooNewsService{
		Fetcher: fetcher.NewYahooNewsRSSFetcher(),
	}
}

func (y *YahooNewsService) FetchAndStore(ctx context.Context) error {
	_, err := y.fetchRSS(ctx)
	return err
}

func (y *YahooNewsService) fetchRSS(ctx context.Context) ([]model.YahooNewsRSS, error) {
	rss := make([]model.YahooNewsRSS, 0, len(conf.YahooNewsMap))
	for _, v := range conf.YahooNewsAllGenre {
		url := conf.YahooNewsMap[v]
		r, err := y.Fetcher.Fetch(ctx, url)
		if err != nil {
			return nil, err
		}
		rss = append(rss, *r)
	}
	return rss, nil
}

func (y *YahooNewsService) scrape() {}
