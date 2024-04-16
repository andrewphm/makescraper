package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type CryptoCurrency struct {
	Name  string
	Price string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	var cryptocurrencies []CryptoCurrency

	selectors := map[string]string{
		"Bitcoin":  "table > tbody > tr:nth-child(1) > td:nth-child(4) > div > a > span",
		"Ethereum": "table > tbody > tr:nth-child(2) > td:nth-child(4) > div > a > span",
	}

	// On visiting the main page, look for the table and the necessary rows using the defined selectors
	c.OnHTML("body", func(e *colly.HTMLElement) {
		for name, selector := range selectors {
			price := e.ChildText(selector)
			cryptocurrencies = append(cryptocurrencies, CryptoCurrency{Name: name, Price: price})
		}
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

	// Print out
	for _, crypto := range cryptocurrencies {
		fmt.Printf("%s Price: %s\n", crypto.Name, crypto.Price)
	}
}
