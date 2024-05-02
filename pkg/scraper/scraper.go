package scraper

import (
	"fmt"
	"net/http"
	"web-scraper/pkg/repository"

	"github.com/PuerkitoBio/goquery"
)

type webScraper struct{}

func NewScraper() repository.Scraper {
	return &webScraper{}
}

func (s *webScraper) Scrape(url string) ([]string, error) {
	// Create HTTP client and request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check HTTP response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error making request, status code: %d", resp.StatusCode)
	}

	// Create a goquery document from the HTTP response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the document to find links
	var results []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			results = append(results, link)
		}
	})
	return results, nil
}
