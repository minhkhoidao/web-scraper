// /cmd/main.go
package main

import (
	"fmt"
	"web-scraper/pkg/di"
)

func main() {
	scraper, err := di.InitializeScraper()
	if err != nil {
		fmt.Printf("Failed to initialize scraper: %v\n", err)
		return
	}
	if scraper == nil {
		fmt.Println("Scraper is nil. Check DI setup.")
		return
	}

	results, err := scraper.Scrape("https://facebook.com")
	if err != nil {
		fmt.Println("Error scraping:", err)
		return
	}
	fmt.Println("Scraped links:", results)
}
