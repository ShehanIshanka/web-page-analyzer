package main

import (
	"log"
	"net/http"

    "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"backend/analyzer"
)

func main() {
	r := gin.Default()

	// For CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/analyze", func(c *gin.Context) {
		var req struct {
			URL string `json:"url"`
		}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid request format",
			})
			return
		}

		resultChan := make(chan struct {
			data       *analyzer.AnalysisResult
			statusCode int
			err        error
		})

		go func() {
			data, statusCode, err := analyzer.AnalyzeURL(req.URL)
			resultChan <- struct {
				data       *analyzer.AnalysisResult
				statusCode int
				err        error
			}{data: data, statusCode: statusCode, err: err}
		}()


		result := <-resultChan

		if result.err != nil {
			c.JSON(result.statusCode, gin.H{
				"success": false,
				"error":   result.err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    result.data,
		})
	})

	
	log.Fatal(r.Run(":8080"))
}
