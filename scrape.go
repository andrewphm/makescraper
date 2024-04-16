package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Define Bitcoin and Ethereum selectors
	bitcoinSelector := "table > tbody > tr:nth-child(1) > td:nth-child(4) > div > a > span"
	ethereumSelector := "table > tbody > tr:nth-child(2) > td:nth-child(4) > div > a > span"

	// On visiting the main page, look for the table and the necessary rows using the defined selectors
	c.OnHTML("body", func(e *colly.HTMLElement) {
		bitcoinPrice := e.ChildText(bitcoinSelector)
		ethereumPrice := e.ChildText(ethereumSelector)

		fmt.Printf("Bitcoin Price: %s\n", bitcoinPrice)
		fmt.Printf("Ethereum Price: %s\n", ethereumPrice)
	})

	// Log errors if there are any
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://coinmarketcap.com/")
}
