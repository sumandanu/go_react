package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sumandanu/go_react/backend/internal/controller"
)

func HomepageHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message":"Welcome to the Tech Company listing API with Golang"})
}


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

	config := cors.DefaultConfig()
    config.AllowAllOrigins = true
    config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour

    router.Use(cors.New(config))

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(rateLimiter())

	// Controllers
	scrapeController := controller.NewScrapeController()

	// Routes
	api := router.Group("/api/v1")
	{
		api.GET("/health", healthCheck)
		api.GET("/news", scrapeController.ScrapeNews)
		api.GET("/newest", scrapeController.ScrapeNews)
		api.GET("/front", scrapeController.ScrapeNews)
		api.GET("/newcomments", scrapeController.ScrapeComments)
		api.GET("/item", scrapeController.ScrapeItemcomments)
		api.GET("/ask", scrapeController.ScrapeAsks)
		api.GET("/show", scrapeController.ScrapeShow)
		api.GET("/jobs", scrapeController.ScrapeJobs)
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
		// "time":   time.Now().Format(time.RFC3339),
	})
}

func rateLimiter() gin.HandlerFunc {
	// Implement rate limiting using redis or in-memory store
	return func(c *gin.Context) {
		// Rate limiting logic
		c.Next()
	}
}