//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package fetcher

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestWebsiteFetcher_Fetch(t *testing.T) {
	type fields struct {
		Client  *http.Client
		Timeout time.Duration
	}
	type args struct {
		ctx context.Context
		URL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "example.com",
			fields: fields{
				Client:  http.DefaultClient,
				Timeout: 1 * time.Second,
			},
			args: args{
				ctx: context.Background(),
				URL: "https://example.com",
			},
			want: trim(exampleDotCom),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &WebsiteFetcher{
				Client:  tt.fields.Client,
				Timeout: tt.fields.Timeout,
			}
			doc, err := y.Fetch(tt.args.ctx, tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebsiteFetcher.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			html, err := doc.Html()
			assert.NoError(t, err)
			got := trim(html)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WebsiteFetcher.Fetch() got diff -got +want %v", diff)
			}
		})
	}
}

func trim(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
}

const exampleDotCom = `<!DOCTYPE html><html><head>
    <title>Example Domain</title>

    <meta charset="utf-8"/>
    <meta http-equiv="Content-type" content="text/html; charset=utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;

    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>
</head>
<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>


</body></html>`
