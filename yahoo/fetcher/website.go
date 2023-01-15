//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package fetcher

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type HTMLDocumentFetcher interface {
	Fetch(ctx context.Context, URL string) (*goquery.Document, error)
}

type WebsiteFetcher struct {
	Client  *http.Client
	Timeout time.Duration
}

func NewWebsiteFetcher() HTMLDocumentFetcher {
	return &WebsiteFetcher{
		Client:  http.DefaultClient,
		Timeout: 1 * time.Second,
	}
}

func (y *WebsiteFetcher) Fetch(ctx context.Context, URL string) (*goquery.Document, error) {
	req, err := http.NewRequest(http.MethodGet, URL, strings.NewReader(""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("accept", "text/html")

	doc, err := y.doRequest(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return doc, nil
}

func (y *WebsiteFetcher) doRequest(ctx context.Context, req *http.Request) (*goquery.Document, error) {
	reqCtx, cancel := context.WithTimeout(ctx, y.Timeout)
	defer cancel()

	reqWithCtx := req.WithContext(reqCtx)
	res, err := y.Client.Do(reqWithCtx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return doc, nil
	}

	switch {
	case res.StatusCode/100 == 4:
		return nil, errors.Errorf("client error occurred. status code = %d", res.StatusCode)
	case res.StatusCode/100 == 5:
		return nil, errors.Errorf("server error occurred. status code = %d", res.StatusCode)
	default:
		return nil, errors.Errorf("unknown occurred. status code = %d", res.StatusCode)
	}
}
