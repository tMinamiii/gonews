package yahoo

const (
	YahooNewsDomestic      = "https://news.yahoo.co.jp/rss/categories/domestic.xml"
	YahooNewsWorld         = "https://news.yahoo.co.jp/rss/categories/world.xml"
	YahooNewsBusiness      = "https://news.yahoo.co.jp/rss/categories/business.xml"
	YahooNewsEntertainment = "https://news.yahoo.co.jp/rss/categories/entertainment.xml"
	YahooNewsSports        = "https://news.yahoo.co.jp/rss/categories/sports.xml"
	YahooNewsIT            = "https://news.yahoo.co.jp/rss/categories/it.xml"
	YahooNewsScience       = "https://news.yahoo.co.jp/rss/categories/science.xml"
	YahooNewsLife          = "https://news.yahoo.co.jp/rss/categories/life.xml"
	YahooNewsLocal         = "https://news.yahoo.co.jp/rss/categories/local.xml"
)

var YahooNewsMap = map[string]string{
	"domestic":      YahooNewsDomestic,
	"world":         YahooNewsWorld,
	"business":      YahooNewsBusiness,
	"entertainment": YahooNewsEntertainment,
	"sports":        YahooNewsSports,
	"it":            YahooNewsIT,
	"science":       YahooNewsScience,
	"life":          YahooNewsLife,
	"local":         YahooNewsLocal,
}
