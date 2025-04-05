package controller

import (
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/sumandanu/go_react/backend/internal/models"
)

type ScrapeController struct {
	mu sync.Mutex
}

func NewScrapeController() *ScrapeController {
	return &ScrapeController{}
}

func (sc *ScrapeController) ScrapeNews(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News
  
	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			Title:      h.ChildText("span.titleline a"), 
			TitleUrl:   h.ChildAttr("span.titleline a", "href"),
			Site:       h.ChildText("span.sitebit.comhead a"), 
			SiteUrl:    h.ChildAttr("span.sitebit.comhead a", "href"),
	  	}
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeNewest(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News
  
	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
	  	}
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			
			if e.Index == 0 {
				news.VoteUrl = e.Attr("href")
			}
			if e.Index == 1 {
				news.Title = e.Text
				news.TitleUrl = e.Attr("href")
			}
			if e.Index == 2 {
				news.Site = e.Text
				news.SiteUrl = e.Attr("href")
			}
		})
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				results[h.Index].PastUrl = e.Attr("href")
			}

			if e.Index == 4 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeFront(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News
  
	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
	  	}
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			
			if e.Index == 0 {
				news.VoteUrl = e.Attr("href")
			}
			if e.Index == 1 {
				news.Title = e.Text
				news.TitleUrl = e.Attr("href")
			}
			if e.Index == 2 {
				news.Site = e.Text
				news.SiteUrl = e.Attr("href")
			}
		})
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				results[h.Index].PastUrl = e.Attr("href")
			}

			if e.Index == 4 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeNewcomments(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.Comment
  
	collector.OnHTML("tr.athing",func(h *colly.HTMLElement) {
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
		results = append(results, comment)
  
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeItemcomments(c *gin.Context) {
	// id := c.Query("id")
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var submission models.Submission
	var results []models.Comment

	collector.OnHTML("table.fatitem",func(h *colly.HTMLElement) {
		submission.Id = h.ChildAttr("tr.athing.submission","id")
		submission.No = h.ChildText("span.rank")
		submission.VoteUrl = h.ChildAttr("a#up_"+ h.ChildAttr("tr.athing.submission","id"), "href")
		submission.Title = h.ChildText("span.titleline > a")
		submission.TitleUrl = h.ChildAttr("span.titleline a", "href")
		submission.Site = h.ChildText("span.sitebit.comhead a")
		submission.SiteUrl = h.ChildAttr("span.sitebit.comhead a", "href")
		submission.Score = h.ChildText("span.score")
	
		h.ForEach("span.subline a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				submission.ByUser = e.Text
				submission.ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				submission.Age = e.Text
				submission.AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				submission.HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				submission.PastUrl = e.Attr("href")
			}

			if e.Index == 4 {
				submission.FavoriteUrl = e.Attr("href")
			}

			if e.Index == 5 {
				submission.Comments = e.Text
				submission.CommentsUrl = e.Attr("href")
			}
		})

		submission.FormAction = h.ChildAttr("form","action")

		h.ForEach("form input", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				submission.FormParentVal = e.Attr("value")
			}

			if e.Index == 1 {
				submission.FormGotoVal = e.Attr("value")
			}

			if e.Index == 2 {
				submission.FormHmacVal = e.Attr("value")
			}
		})
	})

	collector.OnHTML("tr.athing.comtr",func(h *colly.HTMLElement) {
		comment := models.Comment {
			Id:    		h.Attr("id"),
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			ByUser: 	h.ChildText("a.hnuser"),
			ByUrl: 		h.ChildAttr("a.hnuser", "href"),
			Age: 		h.ChildText("span.age a"),
			AgeUrl: 	h.ChildAttr("span.age a", "href"),
			NextUrl: 	h.ChildAttr("span.navs > a", "href"),
			Comment:  	h.ChildText("div.commtext.c00"),
	  	}
		results = append(results, comment)
  
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	// err := collector.Visit("https://news.ycombinator.com/item?id="+id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"submission":submission, "data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeAsks(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News

	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
	  	}
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			
			if e.Index == 0 {
				news.VoteUrl = e.Attr("href")
			}
			if e.Index == 1 {
				news.Title = e.Text
				news.TitleUrl = e.Attr("href")
			}
			if e.Index == 2 {
				news.Site = e.Text
				news.SiteUrl = e.Attr("href")
			}
		})
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	// err := collector.Visit("https://news.ycombinator.com/ask")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeShow(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News
  
	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
	  	}
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			
			if e.Index == 0 {
				news.VoteUrl = e.Attr("href")
			}
			if e.Index == 1 {
				news.Title = e.Text
				news.TitleUrl = e.Attr("href")
			}
			if e.Index == 2 {
				news.Site = e.Text
				news.SiteUrl = e.Attr("href")
			}
		})
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	// err := collector.Visit("https://news.ycombinator.com/front")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}

func (sc *ScrapeController) ScrapeJobs(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	)
  
	var nextLink string
	var results []models.News
  
	collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
			Id:    		h.Attr("id"),
			No: 		h.ChildText("span.rank"), 
			VoteUrl: 	h.ChildAttr("a#up_"+h.Attr("id"), "href"),
			Title:      h.ChildText("span.titleline a"), 
			TitleUrl:   h.ChildAttr("span.titleline a", "href"),
			Site:       h.ChildText("span.sitebit.comhead a"), 
			SiteUrl:    h.ChildAttr("span.sitebit.comhead a", "href"),
	  	}
		
		results = append(results, news)
  
	})
  
	collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		results[h.Index].Score = h.ChildText("span.score")
		h.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if e.Index == 0 {
				results[h.Index].ByUser = e.Text
				results[h.Index].ByUrl = e.Attr("href")
			}

			if e.Index == 1 {
				results[h.Index].Age = e.Text
				results[h.Index].AgeUrl = e.Attr("href")
			}

			if e.Index == 2 {
				results[h.Index].HideUrl = e.Attr("href")
			}

			if e.Index == 3 {
				results[h.Index].Comments = e.Text
				results[h.Index].CommentsUrl = e.Attr("href")
			}
		})
	})

	collector.OnHTML("a.morelink",func(h *colly.HTMLElement) {
		nextLink = h.Attr("href")
	})
  
	err := collector.Visit(strings.Replace(c.Request.URL.String(), "/api/v1", "https://news.ycombinator.com", -1))
	// err := collector.Visit("https://news.ycombinator.com/jobs")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"data": results,  "nextLink" : nextLink })
}