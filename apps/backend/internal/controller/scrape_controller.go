package controller

import (
	"net/http"
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

func (sc *ScrapeController) ScrapeWebsite(c *gin.Context) {
	collector := colly.NewCollector(
		colly.AllowedDomains("news.ycombinator.com"),
	  )
  
	  var results []models.News
  
	  collector.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
		news := models.News {
		  Id:    h.Attr("id"),
		  Title:    h.ChildText("a[rel=nofollow]"),
		  TitleUrl: h.ChildAttr("span.titleline a", "href"),
		  Site:     h.ChildText("span.sitestr"),
		  SiteUrl:  h.ChildAttr("span.sitebit.comhead a", "href"),
		}
		results = append(results, news)
  
	  })
  
	  collector.OnHTML("td.subtext",func(h *colly.HTMLElement) {
		results[h.Index].Score = h.ChildText("span.score")
		results[h.Index].ByUser = h.ChildText("a.hnuser")
		results[h.Index].ByUrl = h.ChildAttr("a.hnuser","href")
		results[h.Index].Age = h.ChildText("span.age a")
		results[h.Index].AgeUrl = h.ChildAttr("span.age a","href")
		results[h.Index].HideUrl = "hide?id="+ results[h.Index].Id +"&goto=newest"
		results[h.Index].PastUrl = h.ChildAttr("a.hnpast","href")
		results[h.Index].CommentsUrl = "item?id="+ results[h.Index].Id
	  })
  
	//   collector.Visit("https://news.ycombinator.com/newest")
  
	//   b, _ := json.Marshal(newsList)
	//   fmt.Print(string(b))
	// url := c.Query("url")
	// if url == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "url parameter is required"})
	// 	return
	// }

	// var results []map[string]string
	// collector := colly.NewCollector(
	// 	colly.Async(true),
	// )

	// collector.OnHTML("article", func(e *colly.HTMLElement) {
	// 	sc.mu.Lock()
	// 	defer sc.mu.Unlock()
		
	// 	results = append(results, map[string]string{
	// 		"title":   e.ChildText("h2"),
	// 		"url":     e.Request.AbsoluteURL(e.Attr("href")),
	// 		"summary": e.ChildText("p.summary"),
	// 	})
	// })

	// collector.OnError(func(r *colly.Response, err error) {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"url":   r.Request.URL,
	// 		"error": err.Error(),
	// 	})
	// })

	err := collector.Visit("https://news.ycombinator.com/newest")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collector.Wait()
	c.JSON(http.StatusOK, gin.H{"results": results})
}