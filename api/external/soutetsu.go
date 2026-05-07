package external

import (
	"fmt"
	"net/http/cookiejar"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const (
	sotetsuLoginPageURL = "https://mypage.sotetsu-point.jp/PISM010_00"
	sotetsuLoginPostURL = "https://mypage.sotetsu-point.jp/PISM010_01"
	sotetsuMyPageURL    = "https://mypage.sotetsu-point.jp/PISM020_00"
)

type SotetsuClient struct {
	jar *cookiejar.Jar
}

type SotetsuData struct {
	Name        string
	Point       int
	PointExpiry string
	Mile        int
	MileExpiry  string
	Rank        string
}

func NewSotetsuClient() (*SotetsuClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &SotetsuClient{jar: jar}, nil
}

func (s *SotetsuClient) newCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.UserAgent(userAgent),
	)
	c.SetCookieJar(s.jar)
	return c
}

// Login はログインページからCSRFトークンを取得し、認証する
func (s *SotetsuClient) Login(email, password string) error {
	c := s.newCollector()
	var loginErr error
	loggedIn := false
	posted := false // 無限ループ防止: ログインPOSTは1回のみ

	c.OnHTML("form#PISM010", func(e *colly.HTMLElement) {
		if posted {
			return
		}
		posted = true
		token := e.ChildAttr("input[name='jp.hitachisoft.message.TOKEN']", "value")
		if err := e.Request.Post(sotetsuLoginPostURL, map[string]string{
			"userId":                       email,
			"passWord":                     password,
			"jp.hitachisoft.message.TOKEN": token,
		}); err != nil {
			loginErr = fmt.Errorf("POSTリクエスト失敗: %w", err)
		}
	})

	// ログイン成功判定: PISM020_00 (マイページ) へ遷移したか
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "PISM020") {
			loggedIn = true
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		loginErr = err
	})

	if err := c.Visit(sotetsuLoginPageURL); err != nil {
		return fmt.Errorf("ログインページアクセス失敗: %w", err)
	}
	if loginErr != nil {
		return loginErr
	}
	if !loggedIn {
		return fmt.Errorf("ログイン失敗: 認証情報を確認してください")
	}
	return nil
}

// FetchAll はマイページから全データを取得する
func (s *SotetsuClient) FetchAll() (*SotetsuData, error) {
	c := s.newCollector()
	data := &SotetsuData{}
	var fetchErr error

	c.OnHTML("main.container", func(e *colly.HTMLElement) {
		// 会員名: "菊池　裕夢 様" → " 様" を除去
		nameRaw := strings.TrimSpace(e.DOM.Find("h1.parts-title03").Text())
		data.Name = strings.TrimSpace(strings.TrimSuffix(nameRaw, "様"))

		// ポイント (color-blue の最初の要素)
		pointStr := strings.TrimSpace(e.DOM.Find(".mypage-status__situation__txt.color-blue").First().Text())
		data.Point, _ = strconv.Atoi(strings.ReplaceAll(pointStr, ",", ""))

		// マイル (color-aqua)
		mileStr := strings.TrimSpace(e.DOM.Find(".mypage-status__situation__txt.color-aqua").First().Text())
		data.Mile, _ = strconv.Atoi(strings.ReplaceAll(mileStr, ",", ""))

		// ローゼンランク (color-blue の最後の要素)
		data.Rank = strings.TrimSpace(e.DOM.Find(".mypage-status__situation__txt.color-blue").Last().Text())

		// 有効期限: strong.color-blue の 1つ目=ポイント、2つ目=マイル
		e.DOM.Find("strong.color-blue").Each(func(i int, sel *goquery.Selection) {
			expiry := strings.TrimSpace(sel.Text())
			expiry = strings.ReplaceAll(expiry, "\u00a0", "") // &nbsp; 除去
			switch i {
			case 0:
				data.PointExpiry = expiry
			case 1:
				data.MileExpiry = expiry
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fetchErr = err
	})

	if err := c.Visit(sotetsuMyPageURL); err != nil {
		return nil, fmt.Errorf("マイページアクセス失敗: %w", err)
	}
	if fetchErr != nil {
		return nil, fetchErr
	}
	return data, nil
}
