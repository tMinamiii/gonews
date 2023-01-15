package yahoo

import (
	"context"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestYahooNewRSSClient_Fetch(t *testing.T) {
	type fields struct {
		Client  *http.Client
		Timeout time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		stubServer func(t *testing.T) *httptest.Server
		fields     fields
		args       args
		want       *YahooNewsRSS
		wantErr    bool
	}{
		{
			name: "RSSからXMLを取得してYahooNesRSS構造体にデータを入れる",
			stubServer: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, _ = w.Write([]byte(stubRes))
				}))
			},
			fields: fields{
				Client:  http.DefaultClient,
				Timeout: 10 * time.Second,
			},
			args: args{
				ctx: context.Background(),
			},
			want: &YahooNewsRSS{
				XMLName: xml.Name{
					Space: "",
					Local: "rss",
				},
				Channel: YahooNewsChannel{
					Language:    "ja",
					Copyright:   "© Yahoo Japan",
					PubDateStr:  "Sun, 15 Jan 2023 09:21:54 GMT",
					Title:       " 国際 - Yahoo!ニュース",
					Link:        "https://news.yahoo.co.jp/categories/world?source=rss",
					Description: "Yahoo! JAPANのニュースに掲載されている記事の最新の見出しを提供しています。",
					Image: YahooNewsImage{
						Title:  "Yahoo!ニュース",
						Link:   "https://news.yahoo.co.jp/",
						URL:    "https://s.yimg.jp/images/news/yjnews_s.gif",
						Width:  101,
						Height: 18,
					},
					YahooNewsItems: []YahooNewsItem{
						{
							Title:       "【コラム】「迫害呼訴人」李在明、韓国の歴代大統領が犯した罪のオンパレード(朝鮮日報日本語版)",
							Link:        "https://news.yahoo.co.jp/articles/c5c24a163a67020d4d778958d3a5339a52d7af2a?source=rss",
							PubDateStr:  "Sun, 15 Jan 2023 09:06:23 GMT",
							Image:       "https://newsatcl-pctr.c.yimg.jp/t/amd-img/20230113-00080153-chosun-000-4-view.jpg?pri=l&w=450&h=282&exp=10800",
							Description: "　「私はベルリン市民です」。ジョン・F・ケネディ米大統領が1963年6月26日、西ベルリンでの講演で発した言葉だ。西ベルリンは東ドイツ領土に囲まれた島のような場所だった。西ベルリン市民は共産主義の暴圧",
						},
						{
							Title:       "ネパールで旅客機墜落　少なくとも４０人死亡(産経新聞)",
							Link:        "https://news.yahoo.co.jp/articles/e0755dc3eef6bccee097844fd0fca77a2e2b52e7?source=rss",
							PubDateStr:  "Sun, 15 Jan 2023 08:48:46 GMT",
							Image:       "https://newsatcl-pctr.c.yimg.jp/t/amd-img/default.jpg?pri=l&w=450&h=450&exp=10800",
							Comments:    "https://news.yahoo.co.jp/articles/e0755dc3eef6bccee097844fd0fca77a2e2b52e7/comments",
							Description: "【シンガポール＝森浩】ネパール中部ポカラで１５日、地元のイエティ航空が運航する旅客機が墜落した。乗客６８人と乗務員４人が搭乗しており、地元航空当局によると、少なくとも４０人の死亡が確認された。 航空",
						},
						{
							Title:       "墜落のネパール機に日本人搭乗情報なし(共同通信)",
							Link:        "https://news.yahoo.co.jp/articles/df46e5b59c442e41480affe17ebab215288a12c1?source=rss",
							PubDateStr:  "Sun, 15 Jan 2023 08:39:52 GMT",
							Image:       "https://newsatcl-pctr.c.yimg.jp/t/amd-img/default.jpg?pri=l&w=450&h=450&exp=10800",
							Comments:    "https://news.yahoo.co.jp/articles/df46e5b59c442e41480affe17ebab215288a12c1/comments",
							Description: "　【ニューデリー共同】在ネパール日本大使館は15日、ネパールで墜落した旅客機に日本人の搭乗情報はないと明らかにした。",
						},
					},
				},
				Version: "2.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := tt.stubServer(t)
			f := &YahooNewRSSClient{
				Client:  tt.fields.Client,
				Timeout: tt.fields.Timeout,
			}

			got, err := f.Fetch(tt.args.ctx, server.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("YahooNewRSSClient.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YahooNewRSSClient.Fetch() got diff -got +want %v", diff)
			}
		})
	}
}

var stubRes = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<rss version="2.0">
  <channel>
    <language>ja</language>
    <copyright>© Yahoo Japan</copyright>
    <pubDate>Sun, 15 Jan 2023 09:21:54 GMT</pubDate>
    <title> 国際 - Yahoo!ニュース</title>
    <link>https://news.yahoo.co.jp/categories/world?source=rss</link>
    <description>Yahoo! JAPANのニュースに掲載されている記事の最新の見出しを提供しています。</description>
    <image>
      <title>Yahoo!ニュース</title>
      <link>https://news.yahoo.co.jp/</link>
      <url>https://s.yimg.jp/images/news/yjnews_s.gif</url>
      <width>101</width>
      <height>18</height>
    </image>
    <item>
      <title>【コラム】「迫害呼訴人」李在明、韓国の歴代大統領が犯した罪のオンパレード(朝鮮日報日本語版)</title>
      <link>https://news.yahoo.co.jp/articles/c5c24a163a67020d4d778958d3a5339a52d7af2a?source=rss</link>
      <pubDate>Sun, 15 Jan 2023 09:06:23 GMT</pubDate>
      <image>https://newsatcl-pctr.c.yimg.jp/t/amd-img/20230113-00080153-chosun-000-4-view.jpg?pri=l&amp;w=450&amp;h=282&amp;exp=10800</image>
      <description>　「私はベルリン市民です」。ジョン・F・ケネディ米大統領が1963年6月26日、西ベルリンでの講演で発した言葉だ。西ベルリンは東ドイツ領土に囲まれた島のような場所だった。西ベルリン市民は共産主義の暴圧</description>
    </item>
    <item>
      <title>ネパールで旅客機墜落　少なくとも４０人死亡(産経新聞)</title>
      <link>https://news.yahoo.co.jp/articles/e0755dc3eef6bccee097844fd0fca77a2e2b52e7?source=rss</link>
      <pubDate>Sun, 15 Jan 2023 08:48:46 GMT</pubDate>
      <image>https://newsatcl-pctr.c.yimg.jp/t/amd-img/default.jpg?pri=l&amp;w=450&amp;h=450&amp;exp=10800</image>
      <comments>https://news.yahoo.co.jp/articles/e0755dc3eef6bccee097844fd0fca77a2e2b52e7/comments</comments>
      <description>【シンガポール＝森浩】ネパール中部ポカラで１５日、地元のイエティ航空が運航する旅客機が墜落した。乗客６８人と乗務員４人が搭乗しており、地元航空当局によると、少なくとも４０人の死亡が確認された。 航空</description>
    </item>
    <item>
      <title>墜落のネパール機に日本人搭乗情報なし(共同通信)</title>
      <link>https://news.yahoo.co.jp/articles/df46e5b59c442e41480affe17ebab215288a12c1?source=rss</link>
      <pubDate>Sun, 15 Jan 2023 08:39:52 GMT</pubDate>
      <image>https://newsatcl-pctr.c.yimg.jp/t/amd-img/default.jpg?pri=l&amp;w=450&amp;h=450&amp;exp=10800</image>
      <comments>https://news.yahoo.co.jp/articles/df46e5b59c442e41480affe17ebab215288a12c1/comments</comments>
      <description>　【ニューデリー共同】在ネパール日本大使館は15日、ネパールで墜落した旅客機に日本人の搭乗情報はないと明らかにした。</description>
    </item>
  </channel>
</rss>
`
