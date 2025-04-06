package scraper

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/sumandanu/go_react/backend/internal/models"
)

type WebScraper struct {
	collector *colly.Collector
}

func NewWebScraper() *WebScraper {
	c := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
		colly.Async(true),
	)

	c.AllowURLRevisit = true
	
	// Set limits
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       1 * time.Second,
	})
	
	return &WebScraper{collector: c}
}

func (s *WebScraper) ScrapeNews(url string) (models.ScrapeResult[[]models.News], error) {
	var result models.ScrapeResult[[]models.News]

	s.collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			Title:      h.ChildText("span.titleline a"), 
			TitleUrl:   h.ChildAttr("span.titleline a", "href"),
			Site:       h.ChildText("span.sitebit.comhead a"), 
			SiteUrl:    h.ChildAttr("span.sitebit.comhead a", "href"),
	  	}
		result.Data = append(result.Data, news)
	})
  
	s.collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		result.Data[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				result.Data[h.Index].ByUser = e.Text
				result.Data[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				result.Data[h.Index].Age = e.Text
				result.Data[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				result.Data[h.Index].HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				result.Data[h.Index].Comments = e.Text
				result.Data[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	s.collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		result.NextLink = h.Attr("href")
	})
	
	// OnError callback
	// s.collector.OnError(func(r *colly.Response, err error) {
	// 	return result, err
	// })
	
	// Start scraping
	err := s.collector.Visit(url)
	if err != nil {
		return result, err
	}
	
	s.collector.Wait()
	return result, nil
}

func (s *WebScraper) ScrapeComments(url string) (models.ScrapeResult[[]models.Comment], error) {
	var result models.ScrapeResult[[]models.Comment]

	s.collector.OnHTML("tr.athing",func(h *colly.HTMLElement) {
		comment := models.Comment {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			ByUser: 	h.ChildText("a.hnuser"),
			ByUrl: 		h.ChildAttr("a.hnuser", "href"),
			Age: 		h.ChildText("span.age a"),
			AgeUrl: 	h.ChildAttr("span.age a", "href"),
			ParentUrl: 	h.ChildAttr("span.navs a","href"),
			ContextUrl: "context?id="+h.Attr("id"),
			OnStory:	h.ChildAttr("span.onstory a", "title"), 
			OnStoryUrl:	h.ChildAttr("span.onstory a", "href"),
			Comment:  	h.ChildText("div.commtext.c00"),
	  	}
		  
		result.Data = append(result.Data, comment)
	})

	s.collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		result.NextLink = h.Attr("href")
	})
	
	// Start scraping
	err := s.collector.Visit(url)
	if err != nil {
		return result, err
	}
	
	s.collector.Wait()
	return result, nil
}

func (s *WebScraper) ScrapeCommentsItems(url string) (models.ScrapeResult[[]models.Comment], error) {
	var result models.ScrapeResult[[]models.Comment]

	s.collector.OnHTML("tr.athing",func(h *colly.HTMLElement) {
		comment := models.Comment {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			ByUser: 	h.ChildText("a.hnuser"),
			ByUrl: 		h.ChildAttr("a.hnuser", "href"),
			Age: 		h.ChildText("span.age a"),
			AgeUrl: 	h.ChildAttr("span.age a", "href"),
			ParentUrl: 	h.ChildAttr("span.navs a","href"),
			ContextUrl: "context?id="+h.Attr("id"),
			OnStory:	h.ChildAttr("span.onstory a", "title"), 
			OnStoryUrl:	h.ChildAttr("span.onstory a", "href"),
			Comment:  	h.ChildText("div.commtext.c00"),
	  	}
		  
		result.Data = append(result.Data, comment)
	})

	s.collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		result.NextLink = h.Attr("href")
	})
	
	// Start scraping
	err := s.collector.Visit(url)
	if err != nil {
		return result, err
	}
	
	s.collector.Wait()
	return result, nil
}
