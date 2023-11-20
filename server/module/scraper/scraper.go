package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

const googleSearchUrl = "https://www.google.com/search?q="

type ScraperResult struct {
	Keyword     string
	TotalResult string
	Html        string
	AdWords     int16
	Links       int16
}

type GoogleScraper struct {
	keyword     string
	totalResult string
	html        string
	adWords     int16
	links       int16
	err         error
}

func NewGoogleScraper(keyword string) *GoogleScraper {
	return &GoogleScraper{keyword: keyword}
}

func (s *GoogleScraper) Start() error {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgent)
		fmt.Println("visiting: ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("responseCode: ", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		s.err = err
	})

	c.OnHTML(totalResultSelector, func(h *colly.HTMLElement) {
		totalText := strings.Split(h.Text, " ")
		totalCount := totalText[1]
		s.totalResult = totalCount
	})

	c.OnHTML(adWordSelector, func(h *colly.HTMLElement) {
		s.adWords++
	})

	c.OnHTML(carouselAdWordSelector, func(h *colly.HTMLElement) {
		s.adWords++
	})

	c.OnHTML(sideAdWordSelector, func(h *colly.HTMLElement) {
		s.adWords++
	})

	c.OnHTML(carouselAdLinkSelector, func(h *colly.HTMLElement) {
		s.links++
	})

	c.OnHTML(itemLinkSelector, func(h *colly.HTMLElement) {
		s.links++
	})

	c.OnHTML("html", func(h *colly.HTMLElement) {
		html := string(h.Response.Body)
		s.html = html
	})

	err := c.Visit(s.getSearchUrl())

	if s.err != nil {
		return s.err
	}

	return err
}

func (s *GoogleScraper) GetResult() ScraperResult {
	return ScraperResult{
		Keyword:     s.keyword,
		TotalResult: s.totalResult,
		Html:        s.html,
		AdWords:     s.adWords,
		Links:       s.links,
	}
}

func (s *GoogleScraper) getSearchUrl() string {
	searchKeyword := url.QueryEscape(s.keyword)
	searchUrl := googleSearchUrl + searchKeyword
	return searchUrl
}
