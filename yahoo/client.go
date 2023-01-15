package yahoo

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type YahooNewsFetcher interface {
	Fetch(ctx context.Context, URL string) (*YahooNewsRSS, error)
}

type YahooNewRSSClient struct {
	Client  *http.Client
	Timeout time.Duration
}

func (f *YahooNewRSSClient) Fetch(ctx context.Context, URL string) (*YahooNewsRSS, error) {
	req, err := http.NewRequest(http.MethodGet, URL, strings.NewReader(""))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Set("accept", "application/xml")

	result, err := f.doRequest(ctx, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	m := &YahooNewsRSS{}
	err = xml.Unmarshal(result, m)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return m, nil
}

func (f *YahooNewRSSClient) doRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	reqCtx, cancel := context.WithTimeout(ctx, f.Timeout)
	defer cancel()

	reqWithCtx := req.WithContext(reqCtx)
	res, err := f.Client.Do(reqWithCtx)
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
		return nil, errors.Errorf("bandit server client error occurred. status code = %d", res.StatusCode)
	case res.StatusCode/100 == 5:
		return nil, errors.Errorf("bandit server server error occurred. status code = %d", res.StatusCode)
	default:
		return nil, errors.Errorf("bandit server error occurred. status code = %d", res.StatusCode)
	}
}
