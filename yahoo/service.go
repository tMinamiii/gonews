package yahoo

import "context"

type YahooNews interface {
	FetchAndStore(ctx context.Context)
}

type YahooNewsService struct {
	Fetcher YahooNewsFetcher
}

func NewYahooNewsService() {}

func (y *YahooNewsService) fetchRSS() (*YahooNewsRSS, error) {
	return nil, nil
}

func (y *YahooNewsService) scrape() {}
