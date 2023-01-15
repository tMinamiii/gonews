package yahoo

import "github.com/PuerkitoBio/goquery"

const YahooNewsCSSSelector = ""
const YahooNewsPaginationNextButton = "#uamods > div.pagination.pagination-noBackground.pagination-noBorder > div > ul > li.pagination_item.pagination_item-next > a"
const YahooNewsContent = "#uamods > div.article_body.highLightSearchTarget > div:nth-child(%d) > p"

type YahooNewsScraper struct {
}

func nextPageURL(doc goquery.Document) string {
	doc.Find(YahooNewsPaginationNextButton)
	return ""
}

func (y *YahooNewsScraper) Content(body []byte) (string, error) {
	return "", nil
}
