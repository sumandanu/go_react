package service

import (
	"github.com/sumandanu/go_react/backend/internal/models"
	"github.com/sumandanu/go_react/backend/internal/scraper"
)

type ScrapeService struct {
	scraper *scraper.WebScraper
}

func NewScrapeService() *ScrapeService {
	return &ScrapeService{
		scraper: scraper.NewWebScraper(),
	}
}

func (s *ScrapeService) GetScrapeNews(url string) (models.ScrapeResult[[]models.News], error) {
	return s.scraper.ScrapeNews(url)
}

func (s *ScrapeService) GetScrapeComments(url string) (models.ScrapeResult[[]models.Comment], error) {
	return s.scraper.ScrapeComments(url)
}

func (s *ScrapeService) GetScrapeCommentsItems(url string) (models.ScrapeResult[[]models.Comment], error) {
	return s.scraper.ScrapeCommentsItems(url)
}