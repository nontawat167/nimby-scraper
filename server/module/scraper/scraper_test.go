package scraper_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/nontawat167/nimby-scraper/server/module/scraper"
	"github.com/stretchr/testify/assert"
)

func TestScraper(t *testing.T) {
	keyword := "samsung flip 5"
	s := scraper.NewGoogleScraper(keyword)

	err := s.Start()
	if err != nil {
		assert.FailNow(t, "failed to scraped")
	}

	result := s.GetResult()

	totalStr := strings.ReplaceAll(result.TotalResult, ",", "")
	_, err = strconv.Atoi(totalStr)
	if err != nil {
		assert.FailNow(t, "total result should be number")
	}

	assert.Greater(t, result.Links, int16(0))
	assert.NotEmpty(t, result.Html)
}
