package conf

type YahooNewsGenre string

var YahooNewsAllGenre = []YahooNewsGenre{
	Domestic,
	World,
	Business,
	Entertainment,
	Sports,
	IT,
	Science,
	Life,
	Local,
}

const (
	Domestic      YahooNewsGenre = "domestic"
	World         YahooNewsGenre = "world"
	Business      YahooNewsGenre = "business"
	Entertainment YahooNewsGenre = "entertainment"
	Sports        YahooNewsGenre = "sports"
	IT            YahooNewsGenre = "it"
	Science       YahooNewsGenre = "science"
	Life          YahooNewsGenre = "life"
	Local         YahooNewsGenre = "local"
)

const (
	DomesticURL      = "https://news.yahoo.co.jp/rss/categories/domestic.xml"
	WorldURL         = "https://news.yahoo.co.jp/rss/categories/world.xml"
	BusinessURL      = "https://news.yahoo.co.jp/rss/categories/business.xml"
	EntertainmentURL = "https://news.yahoo.co.jp/rss/categories/entertainment.xml"
	SportsURL        = "https://news.yahoo.co.jp/rss/categories/sports.xml"
	ITURL            = "https://news.yahoo.co.jp/rss/categories/it.xml"
	ScienceURL       = "https://news.yahoo.co.jp/rss/categories/science.xml"
	LifeURL          = "https://news.yahoo.co.jp/rss/categories/life.xml"
	LocalURL         = "https://news.yahoo.co.jp/rss/categories/local.xml"
)

var YahooNewsMap = map[YahooNewsGenre]string{
	Domestic:      DomesticURL,
	World:         WorldURL,
	Business:      BusinessURL,
	Entertainment: EntertainmentURL,
	Sports:        SportsURL,
	IT:            ITURL,
	Science:       ScienceURL,
	Life:          LifeURL,
	Local:         LocalURL,
}
