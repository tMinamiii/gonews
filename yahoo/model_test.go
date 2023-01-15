package yahoo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestYahooNewsItem_PubDate(t *testing.T) {
	y := &YahooNewsItem{
		PubDateStr: "Sun, 15 Jan 2023 08:39:52 GMT",
	}
	got, err := y.PubDate()
	assert.NoError(t, err)

	loc, err := time.LoadLocation("GMT")
	assert.NoError(t, err)

	want := time.Date(2023, time.January, 15, 8, 39, 52, 0, loc)
	assert.True(t, want.Equal(got))
}
