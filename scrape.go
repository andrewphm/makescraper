package main

import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
	"os"
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
		fmt.Println("Error:", err, "Status Code:", r.StatusCode, "URL:", r.Request.URL)
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

	// Serialize the data to JSON
	jsonData, err := json.MarshalIndent(cryptocurrencies, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling data:", err)
	}

	// Print JSON to stdout
	fmt.Println(string(jsonData))

	// Write JSON to a file
	err = os.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}

}
