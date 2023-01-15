//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package fetcher

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tMinamii/gonews/yahoo/model"
)

type RSSFetcher interface {
	Fetch(ctx context.Context, URL string) (*model.YahooNewsRSS, error)
}

type YahooNewsRSSFetcher struct {
	Client  *http.Client
	Timeout time.Duration
}

func NewYahooNewsRSSFetcher() RSSFetcher {
	return &YahooNewsRSSFetcher{
		Client:  http.DefaultClient,
		Timeout: 1 * time.Second,
	}
}

func (y *YahooNewsRSSFetcher) Fetch(ctx context.Context, URL string) (*model.YahooNewsRSS, error) {
	req, err := http.NewRequest(http.MethodGet, URL, strings.NewReader(""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("accept", "application/xml")

	result, err := y.doRequest(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m := &model.YahooNewsRSS{}
	err = xml.Unmarshal(result, m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m, nil
}

func (y *YahooNewsRSSFetcher) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	reqCtx, cancel := context.WithTimeout(ctx, y.Timeout)
	defer cancel()

	reqWithCtx := req.WithContext(reqCtx)
	res, err := y.Client.Do(reqWithCtx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return bodyBytes, nil
	}

	switch {
	case res.StatusCode/100 == 4:
		return nil, errors.Errorf("client error occurred. status code = %d", res.StatusCode)
	case res.StatusCode/100 == 5:
		return nil, errors.Errorf("server server error occurred. status code = %d", res.StatusCode)
	default:
		return nil, errors.Errorf("unknown error occurred. status code = %d", res.StatusCode)
	}
}
