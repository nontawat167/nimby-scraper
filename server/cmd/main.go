package main

import "github.com/nontawat167/nimby-scraper/server/module/scraper"

func main() {
	gs := scraper.NewGoogleScraper("samsung flip 5")
	err := gs.Start()
	if err != nil {
		panic(err)
	}
}
