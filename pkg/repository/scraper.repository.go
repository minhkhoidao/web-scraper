package repository

type Scraper interface {
	Scrape(url string) ([]string, error)
}
