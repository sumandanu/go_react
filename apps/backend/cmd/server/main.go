package main

import (
	"context"
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/gocolly/colly"
	"github.com/sumandanu/go_react/backend/internal/controller"
	// "github.com/sumandanu/go_react/backend/internal/models"
)



func main() {

    // Configuration
    port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
    }

    // Mode setup
	gin.SetMode(gin.ReleaseMode)

	// Router setup
    router := gin.New()

    // c := colly.NewCollector(
    //   colly.AllowedDomains("news.ycombinator.com"),
    // )

    // var newsList []models.News

    // c.OnHTML("tr.athing.submission",func(h *colly.HTMLElement) {
    //   news := models.News {
    //     Id:    h.Attr("id"),
    //     Title:    h.ChildText("a[rel=nofollow]"),
    //     TitleUrl: h.ChildAttr("span.titleline a", "href"),
    //     Site:     h.ChildText("span.sitestr"),
    //     SiteUrl:  h.ChildAttr("span.sitebit.comhead a", "href"),
    //   }
    //   newsList = append(newsList, news)

    // })

    // c.OnHTML("td.subtext",func(h *colly.HTMLElement) {
    //   newsList[h.Index].Score = h.ChildText("span.score")
    //   newsList[h.Index].ByUser = h.ChildText("a.hnuser")
    //   newsList[h.Index].ByUrl = h.ChildAttr("a.hnuser","href")
    //   newsList[h.Index].Age = h.ChildText("span.age a")
    //   newsList[h.Index].AgeUrl = h.ChildAttr("span.age a","href")
    //   newsList[h.Index].HideUrl = "hide?id="+ newsList[h.Index].Id +"&goto=newest"
    //   newsList[h.Index].PastUrl = h.ChildAttr("a.hnpast","href")
    //   newsList[h.Index].CommentsUrl = "item?id="+ newsList[h.Index].Id
    // })

    // c.Visit("https://news.ycombinator.com/newest")

    // b, _ := json.Marshal(newsList)
    // fmt.Print(string(b))

    // 
    // r := gin.Default()
    // r.GET("/ping", func(c *gin.Context) {
    //   c.JSON(http.StatusOK, gin.H{
    //     "message": "pong",
    //   })
    // })

    // r.GET("/news", func(c *gin.Context) {

    //   c.JSON(http.StatusOK, gin.H{
    //     "message": "pong",
    //   })
    // })
    // r.Run("0.0.0.0:8080")

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(rateLimiter())

	scrapeController := controller.NewScrapeController()

	// Routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.GET("/news", scrapeController.ScrapeWebsite)
		// api.POST("/users", createUser)
		// api.GET("/users/:id", getUser)
	}

    // Server configuration
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

  quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	})
}

	func rateLimiter() gin.HandlerFunc {
		// Implement rate limiting using redis or in-memory store
		return func(c *gin.Context) {
			// Rate limiting logic
			c.Next()
		}
	}
// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Configuration
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	// Router setup
// 	router := gin.New()

// 	// Middleware
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())
// 	router.Use(rateLimiter())

// 	// Routes
// 	api := router.Group("/api/v1")
// 	{
// 		api.GET("/health", healthCheck)
// 		// api.POST("/users", createUser)
// 		// api.GET("/users/:id", getUser)
// 	}

// 	// Server configuration
// 	srv := &http.Server{
// 		Addr:         ":" + port,
// 		Handler:      router,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		IdleTimeout:  15 * time.Second,
// 	}

// 	// Graceful shutdown
// 	go func() {
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s\n", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 	<-quit
// 	log.Println("Shutting down server...")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatal("Server forced to shutdown:", err)
// 	}
// 	log.Println("Server exiting")
// }

// func healthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "ok",
// 		"time":   time.Now().Format(time.RFC3339),
// 	})
// }

// func rateLimiter() gin.HandlerFunc {
// 	// Implement rate limiting using redis or in-memory store
// 	return func(c *gin.Context) {
// 		// Rate limiting logic
// 		c.Next()
// 	}
// }