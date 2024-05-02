package di

import (
	"web-scraper/pkg/repository"
	"web-scraper/pkg/scraper"

	"github.com/google/wire"
)

// InitializeScraper is modified to return an error if something goes wrong during the setup.
func InitializeScraper() (repository.Scraper, error) {
	wire.Build(scraper.NewScraper)
	return nil, nil // The second nil is for error, Wire will replace this at runtime
}
